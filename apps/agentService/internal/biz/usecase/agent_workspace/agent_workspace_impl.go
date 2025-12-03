package agent_workspace

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/entity"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/repo"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/convert"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/data/db/model"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/pkg/errors/code"
)

type agentWorkspaceUsecase struct {
	workspaceRepo repo.AgentWorkspaceRepo
	agentRepo     repo.AgentRepo
	log           *log.Helper
}

func NewAgentWorkspaceUsecase(
	workspaceRepo repo.AgentWorkspaceRepo,
	agentRepo repo.AgentRepo,
	logger log.Logger,
) AgentWorkspaceUsecase {
	return &agentWorkspaceUsecase{
		workspaceRepo: workspaceRepo,
		agentRepo:     agentRepo,
		log:           log.NewHelper(logger),
	}
}

func (u *agentWorkspaceUsecase) GetWorkspaceAgents(ctx context.Context, workspaceId string) ([]*entity.Agent, error) {
	// 获取工作区下的所有Agent ID
	agentIDs, err := u.workspaceRepo.GetWorkspaceAgentIDs(ctx, workspaceId)
	if err != nil {
		u.log.Errorf("GetWorkspaceAgents.GetWorkspaceAgentIDs error: %v", err)
		return nil, code.ErrAgentUnKnown
	}

	if len(agentIDs) == 0 {
		return []*entity.Agent{}, nil
	}

	// 获取所有Agent的详细信息
	agents := make([]*entity.Agent, 0, len(agentIDs))
	for _, agentID := range agentIDs {
		// 这里假设Agent表有一个字段可以通过agentID查询
		// 需要在AgentRepo中添加GetByAgentID方法
		agentModel, err := u.agentRepo.GetByAgentID(ctx, agentID)
		if err != nil {
			u.log.Warnf("GetWorkspaceAgents.GetByAgentID error for agentID %s: %v", agentID, err)
			continue
		}
		if agentModel != nil {
			agents = append(agents, convert.AgentPo2Do(agentModel))
		}
	}

	return agents, nil
}

func (u *agentWorkspaceUsecase) AddAgentToWorkspace(ctx context.Context, uid, workspaceId, agentId string) error {
	// 检查Agent是否已经在工作区中
	exists, err := u.workspaceRepo.CheckAgentInWorkspace(ctx, workspaceId, agentId)
	if err != nil {
		u.log.Errorf("AddAgentToWorkspace.CheckAgentInWorkspace error: %v", err)
		return code.ErrAgentUnKnown
	}

	if exists {
		u.log.Warnf("Agent %s already exists in workspace %s", agentId, workspaceId)
		return nil // 已存在，直接返回成功
	}

	// 添加到工作区
	now := time.Now()
	workspaceAgent := &model.AgentWorkspace{
		WorkspaceID: workspaceId,
		AgentID:     agentId,
		UID:         uid,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	err = u.workspaceRepo.AddAgentToWorkspace(ctx, workspaceAgent)
	if err != nil {
		u.log.Errorf("AddAgentToWorkspace.AddAgentToWorkspace error: %v", err)
		return code.ErrAgentUnKnown
	}

	return nil
}

func (u *agentWorkspaceUsecase) RemoveAgentFromWorkspace(ctx context.Context, uid, workspaceId, agentId string) error {
	err := u.workspaceRepo.RemoveAgentFromWorkspace(ctx, workspaceId, agentId)
	if err != nil {
		u.log.Errorf("RemoveAgentFromWorkspace.RemoveAgentFromWorkspace error: %v", err)
		return code.ErrAgentUnKnown
	}

	return nil
}

func (u *agentWorkspaceUsecase) UpdateAgentModelConfig(ctx context.Context, workspace *entity.AgentWorkspace) error {
	// 检查Agent是否在指定的工作区中
	exists, err := u.workspaceRepo.CheckAgentInWorkspace(ctx, workspace.WorkspaceID, workspace.AgentID)
	if err != nil {
		u.log.Errorf("UpdateAgentModelConfig.CheckAgentInWorkspace error: %v", err)
		return code.ErrAgentUnKnown
	}

	if !exists {
		u.log.Warnf("Agent %s not found in workspace %s", workspace.AgentID, workspace.WorkspaceID)
		return code.ErrAgentNotFound
	}

	// 转换为模型对象并更新
	model := convert.AgentWorkspaceDo2Po(workspace)
	err = u.workspaceRepo.UpdateAgentModelConfig(ctx, model)
	if err != nil {
		u.log.Errorf("UpdateAgentModelConfig.UpdateAgentModelConfig error: %v", err)
		return code.ErrAgentUnKnown
	}

	return nil
}
