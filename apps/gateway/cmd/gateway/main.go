package main

//
//import (
//	"context"
//	"flag"
//	"github.com/go-kratos/kratos/v2/config"
//	"github.com/go-kratos/kratos/v2/config/file"
//	"github.com/go-kratos/kratos/v2/middleware/recovery"
//	khttp "github.com/go-kratos/kratos/v2/transport/http"
//	"github.com/johnsonoklii/agentgo/apps/gateway/internal/conf"
//	"github.com/johnsonoklii/agentgo/apps/gateway/internal/discovery"
//	"github.com/johnsonoklii/agentgo/apps/gateway/internal/middleware"
//	"github.com/johnsonoklii/agentgo/apps/gateway/internal/proxy"
//	"github.com/johnsonoklii/agentgo/apps/gateway/internal/service"
//	"github.com/johnsonoklii/agentgo/pkg/errors"
//	"github.com/johnsonoklii/agentgo/pkg/jwt"
//	"github.com/johnsonoklii/agentgo/pkg/redisx"
//
//	"log"
//	"net/http"
//	"time"
//)
//
//var flagconf string
//
//func init() {
//	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
//}
//
//// 全局中间件注册函数
//func ChainMiddleware(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
//	for _, m := range middlewares {
//		handler = m(handler)
//	}
//	return handler
//}
//
//func main() {
//
//	flag.Parse()
//	c := config.New(
//		config.WithSource(
//			file.NewSource(flagconf),
//		),
//	)
//	defer c.Close()
//
//	if err := c.Load(); err != nil {
//		panic(err)
//	}
//
//	var bc conf.Bootstrap
//	if err := c.Scan(&bc); err != nil {
//		panic(err)
//	}
//
//	jwt.InitSecret(bc.Auth.JwtSecret, time.Duration(bc.Auth.Ttl)*time.Second)
//
//	jwt.InitTokenRedis(&redisx.Config{
//		Addr:     bc.TokenRedis.Addr,
//		Password: bc.TokenRedis.Password,
//	})
//
//	consul, err := discovery.NewConsul("127.0.0.1:8500")
//	if err != nil {
//		log.Fatalf("Consul init failed: %v", err)
//	}
//
//	proxyHandler := proxy.NewProxyHandler(consul)
//
//	var opts = []khttp.ServerOption{
//		khttp.Middleware(
//			recovery.Recovery(),
//			middleware.JWTMiddleware,
//		),
//		khttp.Address(bc.Server.Addr),
//		khttp.ErrorEncoder(errors.CustomErrorEncoder),
//	}
//
//	svr := khttp.NewServer(opts...)
//
//	svr.HandlePrefix("/", http.HandlerFunc(service.LogoutHandler))
//	proxy.Proxy_HTTP_Handler(proxyHandler)
//
//	//svr.Handle("/", proxyHandler)
//	//svr.Handle("/v1/auth/logout", http.HandlerFunc(service.LogoutHandler))
//
//	//所有 HTTP 方法
//	methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"}
//
//	r := svr.Route("/")
//	//for _, m := range methods {
//	//	r.Handle(m, "/v1/auth/logout", proxy.Proxy_HTTP_Handler(proxyHandler))
//	//}
//
//	for _, m := range methods {
//		r.Handle(m, "/*", proxy.Proxy_HTTP_Handler(proxyHandler))
//	}
//
//	err = svr.Start(context.Background())
//	if err != nil {
//		panic(err)
//	}
//}
