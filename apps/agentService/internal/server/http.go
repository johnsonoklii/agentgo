package server

import (
	v1 "github.com/johnsonoklii/agentgo/apps/agentService/api/agent/v1"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/conf"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/service"
	"github.com/johnsonoklii/agentgo/pkg/jwt"
	"github.com/johnsonoklii/agentgo/pkg/middleware"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, agentService *service.AgentService, jwtManger *jwt.JWTManager, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			middleware.RedirectMiddleware(jwtManger, log.NewHelper(logger)),
			middleware.AuthMiddleware(jwtManger, log.NewHelper(logger)),
		),
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
	v1.RegisterAgentHTTPServer(srv, agentService)
	return srv
}
