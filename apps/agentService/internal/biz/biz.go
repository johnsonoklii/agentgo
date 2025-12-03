package biz

import (
	"github.com/google/wire"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/usecase/agent"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/usecase/agent_workspace"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/usecase/modal"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/usecase/provider"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	agent.NewAgentUsecase,
	agent.NewAgentVersionUsecase,
	provider.NewProviderUsecase,
	modal.NewModalUsecase,
	agent_workspace.NewAgentWorkspaceUsecase,
)
