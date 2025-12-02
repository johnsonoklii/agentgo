package repo

import (
	"context"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/data/db/model"
)

type AgentWorkspaceRepo interface {
	// AddAgentToWorkspace 添加Agent到工作区
	AddAgentToWorkspace(ctx context.Context, workspaceAgent *model.AgentWorkspace) error
	// RemoveAgentFromWorkspace 从工作区移除Agent
	RemoveAgentFromWorkspace(ctx context.Context, workspaceId, agentId string) error
	// GetWorkspaceAgents 获取工作区下的所有Agent ID列表
	GetWorkspaceAgentIDs(ctx context.Context, workspaceId string) ([]string, error)
	// CheckAgentInWorkspace 检查Agent是否在工作区中
	CheckAgentInWorkspace(ctx context.Context, workspaceId, agentId string) (bool, error)
}
