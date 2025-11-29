package agent

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/entity"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/repo"
)

type agentUsecase struct {
	agentRepo repo.AgentRepo
	log       *log.Helper
}

func NewAgentUsecase(agentRepo repo.AgentRepo, logger log.Logger) AgentUsecase {
	return &agentUsecase{
		agentRepo: agentRepo,
		log:       log.NewHelper(logger),
	}
}

func (u *agentUsecase) Create(ctx context.Context, req *CreateAgentRequest) error {
	return nil
}
func (u *agentUsecase) GetAgentList(ctx context.Context, userId int64) ([]*entity.Agent, error) {
	return nil, nil
}
func (u *agentUsecase) GetAgent(ctx context.Context, agentId int64) (*entity.Agent, error) {
	return nil, nil
}
func (u *agentUsecase) Update(ctx context.Context, req *UpdateUserRequest) error {
	return nil
}
func (u *agentUsecase) Delete(ctx context.Context, agentId int64) error {
	return nil
}
