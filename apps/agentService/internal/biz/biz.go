package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/usecase/agent"
	"github.com/johnsonoklii/agentgo/pkg/jwt"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(agent.NewAgentUsecase, NewJwtManager)

func NewJwtManager(option *jwt.Options, logger log.Logger) *jwt.JWTManager {
	return jwt.NewJWTManager(option, logger)
}
