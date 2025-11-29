package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/johnsonoklii/agentgo/apps/gateway/internal/conf"
	"github.com/johnsonoklii/agentgo/apps/gateway/internal/data"
	"github.com/johnsonoklii/agentgo/apps/gateway/internal/discovery"
	jwt2 "github.com/johnsonoklii/agentgo/apps/gateway/internal/middleware/jwt"
	"github.com/johnsonoklii/agentgo/apps/gateway/internal/proxy"
	"github.com/johnsonoklii/agentgo/pkg/jwt"
	"log"
	"net/http"
	"time"
)

var flagconf string

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

// 全局中间件注册函数
func ChainMiddleware(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middlewares {
		handler = m(handler)
	}
	return handler
}

func main() {
	flag.Parse()
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	// jwt
	jwt.InitSecret(bc.Auth.Jwt.Secret, time.Duration(bc.Auth.Jwt.Expire)*time.Second)

	// db
	_, err := data.Init(bc.Data)
	if err != nil {
		panic(err)
	}

	// consul
	consul, err := discovery.NewConsul(bc.Registry.Consul.Address)
	if err != nil {
		log.Fatalf("Consul init failed: %v", err)
	}

	proxyHandler := proxy.NewProxyHandler(consul)
	mux := http.NewServeMux()

	mux.Handle("/", proxyHandler)

	// 注册全局中间件
	handler := ChainMiddleware(mux, jwt2.JWTMiddleware, jwt2.RedirectMiddleware) //middleware.JWTMiddleware,

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Gateway started at :8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
