package repo

import (
	"context"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/data/db/model"
)

type AgentRepo interface {
	Create(ctx context.Context, agent *model.Agent) error
	Update(ctx context.Context, agentId int64, updates map[string]interface{}) error
	Get(ctx context.Context, agentId int64) (*model.Agent, error)
	GetByAgentID(ctx context.Context, agentID string) (*model.Agent, error)
	List(ctx context.Context, userId int64) ([]*model.Agent, error)
	Delete(ctx context.Context, agentId int64) error
}
