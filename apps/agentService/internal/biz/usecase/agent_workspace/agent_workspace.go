package agent_workspace

import (
	"context"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/entity"
)

type AgentWorkspaceUsecase interface {
	// GetWorkspaceAgents 获取工作区下的所有Agent
	GetWorkspaceAgents(ctx context.Context, workspaceId string) ([]*entity.Agent, error)
	// AddAgentToWorkspace 添加Agent到工作区
	AddAgentToWorkspace(ctx context.Context, uid, workspaceId, agentId string) error
	// RemoveAgentFromWorkspace 从工作区移除Agent
	RemoveAgentFromWorkspace(ctx context.Context, uid, workspaceId, agentId string) error
}
