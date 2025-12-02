package repo

import (
	"context"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/data/db/model"
)

type AgentVersionRepo interface {
	Create(ctx context.Context, version *model.AgentVersion) error
	GetByID(ctx context.Context, id string) (*model.AgentVersion, error)
	GetByAgentIDAndVersion(ctx context.Context, agentID, versionNumber string) (*model.AgentVersion, error)
	ListByAgentID(ctx context.Context, agentID string) ([]*model.AgentVersion, error)
	Update(ctx context.Context, id string, updates map[string]interface{}) error
}
