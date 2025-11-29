package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	authV1 "github.com/johnsonoklii/agentgo/apps/baseService/api/auth/v1"
	userV1 "github.com/johnsonoklii/agentgo/apps/baseService/api/user/v1"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/conf"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/middleware"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/service"
	"github.com/johnsonoklii/agentgo/pkg/errors"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, userService *service.UserService, authService *service.AuthService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			middleware.UserMiddleware,
		),
		http.ErrorEncoder(errors.CustomErrorEncoder),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	srv := http.NewServer(opts...)
	userV1.RegisterUserHTTPServer(srv, userService)
	authV1.RegisterAuthHTTPServer(srv, authService)
	return srv
}
