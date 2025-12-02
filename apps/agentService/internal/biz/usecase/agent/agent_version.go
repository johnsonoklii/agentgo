package agent

import (
	"context"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/entity"
)

type PublishAgentRequest struct {
	AgentID       string
	VersionNumber string
	ChangeLog     string
	UserID        string
}

type AgentVersionUsecase interface {
	PublishAgent(ctx context.Context, req *PublishAgentRequest) (*entity.AgentVersion, error)
	GetAgentVersions(ctx context.Context, agentID string) ([]*entity.AgentVersion, error)
	GetAgentVersion(ctx context.Context, agentID, versionNumber string) (*entity.AgentVersion, error)
	GetAgentLatestVersion(ctx context.Context, agentID string) (*entity.AgentVersion, error)
}
