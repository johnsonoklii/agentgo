package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/johnsonoklii/agentgo/apps/agentService/api/agent/v1"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/conf"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/service"
	"github.com/johnsonoklii/agentgo/pkg/errors"
	"github.com/johnsonoklii/agentgo/pkg/middleware"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, agentService *service.AgentService) *http.Server {
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
	v1.RegisterAgentHTTPServer(srv, agentService)
	return srv
}
