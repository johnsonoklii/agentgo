package db

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/repo"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/data/db/model"
	"gorm.io/gorm"
)

type agentVersionRepo struct {
	data *Data
	log  *log.Helper
}

func NewAgentVersionRepo(data *Data, logger log.Logger) repo.AgentVersionRepo {
	return &agentVersionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *agentVersionRepo) Create(ctx context.Context, version *model.AgentVersion) error {
	return r.data.DB.WithContext(ctx).Create(version).Error
}

func (r *agentVersionRepo) GetByID(ctx context.Context, id string) (*model.AgentVersion, error) {
	version := &model.AgentVersion{}
	err := r.data.DB.WithContext(ctx).Where("id = ?", id).First(version).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return version, nil
}

func (r *agentVersionRepo) GetByAgentIDAndVersion(ctx context.Context, agentID, versionNumber string) (*model.AgentVersion, error) {
	version := &model.AgentVersion{}
	err := r.data.DB.WithContext(ctx).
		Where("agent_id = ? AND version_number = ?", agentID, versionNumber).
		First(version).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return version, nil
}

func (r *agentVersionRepo) ListByAgentID(ctx context.Context, agentID string) ([]*model.AgentVersion, error) {
	var versions []*model.AgentVersion
	err := r.data.DB.WithContext(ctx).
		Where("agent_id = ? AND deleted_at IS NULL", agentID).
		Order("created_at DESC").
		Find(&versions).Error
	return versions, err
}

func (r *agentVersionRepo) Update(ctx context.Context, id string, updates map[string]interface{}) error {
	if _, ok := updates["updated_at"]; !ok {
		updates["updated_at"] = time.Now()
	}
	return r.data.DB.WithContext(ctx).
		Model(&model.AgentVersion{}).
		Where("id = ?", id).
		Updates(updates).Error
}
