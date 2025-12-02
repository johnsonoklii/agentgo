package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/repo"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/data/db/model"
	"gorm.io/gorm"
	"time"
)

type agentRepo struct {
	data *Data
	log  *log.Helper
}

func NewAgentRepo(data *Data, logger log.Logger) repo.AgentRepo {
	return &agentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *agentRepo) Create(ctx context.Context, agent *model.Agent) error {
	return r.data.DB.WithContext(ctx).Create(agent).Error
}

func (r *agentRepo) Update(ctx context.Context, agentId int64, updates map[string]interface{}) error {
	if _, ok := updates["updated_at"]; !ok {
		updates["updated_at"] = time.Now()
	}

	return r.data.DB.WithContext(ctx).Model(&model.Agent{}).Where("id = ?", agentId).Updates(updates).Error
}

func (r *agentRepo) Get(ctx context.Context, agentId int64) (*model.Agent, error) {
	agent := &model.Agent{}
	err := r.data.DB.WithContext(ctx).Where("id = ?", agentId).First(agent).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return agent, nil
}

func (r *agentRepo) GetByAgentID(ctx context.Context, agentID string) (*model.Agent, error) {
	agent := &model.Agent{}
	err := r.data.DB.WithContext(ctx).Where("agent_id = ?", agentID).First(agent).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return agent, nil
}

func (r *agentRepo) List(ctx context.Context, userId int64) ([]*model.Agent, error) {
	agentList := []*model.Agent{}
	err := r.data.DB.WithContext(ctx).Where("user_id = ?", userId).Find(&agentList).Error
	return agentList, err
}

func (r *agentRepo) Delete(ctx context.Context, agentId int64) error {
	return r.data.DB.WithContext(ctx).Where("id = ?", agentId).Delete(&model.Agent{}).Error
}
