package biz

import (
	"github.com/google/wire"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/usecase/agent"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	agent.NewAgentUsecase,
	agent.NewAgentVersionUsecase,
)
