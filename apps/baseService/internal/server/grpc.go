package server

import (
	authV1 "github.com/johnsonoklii/agentgo/apps/baseService/api/auth/v1"
	userV1 "github.com/johnsonoklii/agentgo/apps/baseService/api/user/v1"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/conf"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, userService *service.UserService, authService *service.AuthService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}

	srv := grpc.NewServer(opts...)
	userV1.RegisterUserServer(srv, userService)
	authV1.RegisterAuthServer(srv, authService)

	return srv
}
