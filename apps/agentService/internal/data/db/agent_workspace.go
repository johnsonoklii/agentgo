package db

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/repo"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/data/db/model"
	"gorm.io/gorm"
)

type agentWorkspaceRepo struct {
	data *Data
	log  *log.Helper
}

func NewAgentWorkspaceRepo(data *Data, logger log.Logger) repo.AgentWorkspaceRepo {
	return &agentWorkspaceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *agentWorkspaceRepo) AddAgentToWorkspace(ctx context.Context, workspaceAgent *model.AgentWorkspace) error {
	return r.data.DB.WithContext(ctx).Create(workspaceAgent).Error
}

func (r *agentWorkspaceRepo) RemoveAgentFromWorkspace(ctx context.Context, workspaceId, agentId string) error {
	return r.data.DB.WithContext(ctx).
		Where("workspace_id = ? AND agent_id = ?", workspaceId, agentId).
		Delete(&model.AgentWorkspace{}).Error
}

func (r *agentWorkspaceRepo) GetWorkspaceAgentIDs(ctx context.Context, workspaceId string) ([]string, error) {
	var agentIDs []string
	err := r.data.DB.WithContext(ctx).
		Model(&model.AgentWorkspace{}).
		Where("workspace_id = ? AND deleted_at IS NULL", workspaceId).
		Pluck("agent_id", &agentIDs).Error
	
	if err != nil {
		return nil, err
	}
	
	return agentIDs, nil
}

func (r *agentWorkspaceRepo) CheckAgentInWorkspace(ctx context.Context, workspaceId, agentId string) (bool, error) {
	var count int64
	err := r.data.DB.WithContext(ctx).
		Model(&model.AgentWorkspace{}).
		Where("workspace_id = ? AND agent_id = ? AND deleted_at IS NULL", workspaceId, agentId).
		Count(&count).Error
	
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	
	return count > 0, nil
}
