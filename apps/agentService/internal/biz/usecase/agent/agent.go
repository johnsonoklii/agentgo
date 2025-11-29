package agent

import (
	"context"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/entity"
)

type CreateAgentRequest struct {
	Name             string
	Avatar           string
	Description      string
	SystemPrompt     string
	ToolIds          []string
	KnowledgeBaseIds []string
	PublishedVersion string
	Enabled          bool
	UserId           int64
	UserName         string
	ToolPresetParams map[string]map[string]map[string]string
	MultiModel       bool
}

type UpdateUserRequest struct {
	AgentId          int64
	Name             string
	Avatar           string
	Description      string
	SystemPrompt     string
	ToolIds          []string
	KnowledgeBaseIds []string
	PublishedVersion string
	Enabled          bool
	ToolPresetParams map[string]map[string]map[string]string
	MultiModel       bool
}

type AgentUsecase interface {
	Create(ctx context.Context, req *CreateAgentRequest) error
	GetAgentList(ctx context.Context, userId int64) ([]*entity.Agent, error)
	GetAgent(ctx context.Context, agentId int64) (*entity.Agent, error)
	Update(ctx context.Context, req *UpdateUserRequest) error
	Delete(ctx context.Context, agentId int64) error
}
