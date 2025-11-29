package main

import (
	"flag"
	"github.com/johnsonoklii/agentgo/pkg/jwt"
	"os"
	"time"

	"github.com/johnsonoklii/agentgo/apps/agentService/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	_ "go.uber.org/automaxprocs"
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

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
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

	Name = bc.Server.Grpc.Name
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

	// consul
	var rc conf.Registry
	if err := c.Scan(&rc); err != nil {
		panic(err)
	}

	// jwt
	var auth conf.Auth
	if err := c.Scan(&auth); err != nil {
		panic(err)
	}
	jwtOption := jwt.Options{
		SecretKey: []byte(auth.Jwt.Secret),
		Issuer:    auth.Jwt.Issuer,
		Expire:    time.Duration(auth.Jwt.Expire) * time.Second,

		RdsAddr: auth.Redis.Addr,
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, &rc, &jwtOption, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
