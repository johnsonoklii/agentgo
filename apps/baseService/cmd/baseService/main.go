package main

import (
	"context"
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/conf"
	"github.com/johnsonoklii/agentgo/pkg/jwt"
	"github.com/johnsonoklii/agentgo/pkg/utils"
	_ "go.uber.org/automaxprocs"
	"os"
	"time"
)

// go build -ldflags "-X main.Version=x.y.z"
var (

	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, c *conf.Server, gs *grpc.Server, hs *http.Server, rr registry.Registrar) *kratos.App {
	ip, err := utils.GetLocalIP()
	if err != nil {
		panic(err)
	}

	grpcIp := ip + c.Grpc.Addr
	httpIp := ip + c.Http.Addr

	// 创建带有元数据的服务实例用于gRPC
	grpcInstance := &registry.ServiceInstance{
		ID:      id + "-base-grpc",
		Name:    Name + "-grpc",
		Version: Version,
		Metadata: map[string]string{
			"protocol": "grpc",
		},
		Endpoints: []string{"grpc://" + grpcIp},
	}

	//创建带有元数据的服务实例用于HTTP
	httpInstance := &registry.ServiceInstance{
		ID:      id + "-base-http",
		Name:    Name + "-http",
		Version: Version,
		Metadata: map[string]string{
			"protocol": "http",
		},
		Endpoints: []string{"http://" + httpIp},
	}

	//注册两个服务实例
	if err := rr.Register(context.Background(), grpcInstance); err != nil {
		logger.Log(log.LevelError, "msg", "register grpc service failed", "error", err)
	}

	if err := rr.Register(context.Background(), httpInstance); err != nil {
		logger.Log(log.LevelError, "msg", "register http service failed", "error", err)
	}

	// 创建 Kratos App，同时启动两个 Server
	return kratos.New(
		kratos.ID(id+"-"+Name),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Logger(logger),
		kratos.Server(hs, gs),
	)
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

	Name = bc.Server.Name
	Version = bc.Server.Version

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"gateway.id", id,
		"gateway.name", Name,
		"gateway.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	jwt.InitSecret(bc.Auth.Jwt.Secret, time.Duration(bc.Auth.Jwt.Expire)*time.Second)

	app, cleanup, err := wireApp(bc.Server, bc.Data, bc.Registry, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
