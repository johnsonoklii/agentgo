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
	// UpdateAgentModelConfig 设置Agent的模型配置
	UpdateAgentModelConfig(ctx context.Context, workspace *entity.AgentWorkspace) error
}

// AgentModelConfig Agent模型配置实体
type AgentModelConfig struct {
	Provider    string  // 模型提供商
	Model       string  // 模型名称
	Temperature float32 // 温度参数
	MaxTokens   int32   // 最大token数
	ApiKey      string  // API密钥
}
