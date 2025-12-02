package agent

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/entity"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/repo"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/convert"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/data/db/model"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/pkg/errors/code"
	"github.com/johnsonoklii/agentgo/pkg/utils"
	"time"
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

func (u *agentUsecase) Create(ctx context.Context, uid string, req *CreateAgentRequest) (*entity.Agent, error) {
	now := time.Now()
	agentModel := model.Agent{
		AgentID:          utils.NewUUID(),
		Name:             req.Name,
		Avatar:           req.Avatar,
		Description:      req.Description,
		SystemPrompt:     req.SystemPrompt,
		ToolIds:          req.ToolIds,
		KnowledgeBaseIds: req.KnowledgeBaseIds,
		UID:              uid,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	err := u.agentRepo.Create(ctx, &agentModel)
	if err != nil {
		u.log.Errorf("agentUsecase.Create error: %v", err)
		return nil, code.ErrAgentUnKnown
	}
	return convert.AgentPo2Do(&agentModel), err
}

func (u *agentUsecase) GetAgentList(ctx context.Context, userId int64) ([]*entity.Agent, error) {
	agentModels, err := u.agentRepo.List(ctx, userId)
	if err != nil {
		u.log.Errorf("agentUsecase.GetAgentList error: %v", err)
		return nil, code.ErrAgentUnKnown
	}

	agents := make([]*entity.Agent, 0, len(agentModels))
	for _, agentModel := range agentModels {
		agents = append(agents, convert.AgentPo2Do(agentModel))
	}

	return agents, nil
}

func (u *agentUsecase) GetAgent(ctx context.Context, agentId int64) (*entity.Agent, error) {
	agentModel, err := u.agentRepo.Get(ctx, agentId)
	if err != nil {
		u.log.Errorf("agentUsecase.GetAgent error: %v", err)
		return nil, code.ErrAgentUnKnown
	}

	if agentModel == nil {
		return nil, code.ErrAgentNotFound
	}

	return convert.AgentPo2Do(agentModel), nil
}

func (u *agentUsecase) Update(ctx context.Context, req *UpdateUserRequest) error {
	updates := make(map[string]interface{})

	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Avatar != "" {
		updates["avatar"] = req.Avatar
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.SystemPrompt != "" {
		updates["system_prompt"] = req.SystemPrompt
	}
	if req.ToolIds != nil {
		updates["tool_ids"] = req.ToolIds
	}
	if req.KnowledgeBaseIds != nil {
		updates["knowledge_base_ids"] = req.KnowledgeBaseIds
	}
	if req.PublishedVersion != "" {
		updates["published_version"] = req.PublishedVersion
	}
	updates["enabled"] = req.Enabled
	if req.ToolPresetParams != nil {
		updates["tool_preset_params"] = req.ToolPresetParams
	}
	updates["multi_model"] = req.MultiModel

	err := u.agentRepo.Update(ctx, req.AgentId, updates)
	if err != nil {
		u.log.Errorf("agentUsecase.Update error: %v", err)
		return code.ErrAgentUnKnown
	}

	return nil
}

func (u *agentUsecase) Delete(ctx context.Context, agentId int64) error {
	err := u.agentRepo.Delete(ctx, agentId)
	if err != nil {
		u.log.Errorf("agentUsecase.Delete error: %v", err)
		return code.ErrAgentUnKnown
	}

	return nil
}

func (u *agentUsecase) ToggleStatus(ctx context.Context, agentID string, enabled bool) (*entity.Agent, error) {
	// 1. 获取 Agent
	agentModel, err := u.agentRepo.GetByAgentID(ctx, agentID)
	if err != nil {
		u.log.Errorf("agentUsecase.ToggleStatus.GetByAgentID error: %v", err)
		return nil, code.ErrAgentUnKnown
	}
	if agentModel == nil {
		return nil, code.ErrAgentNotFound
	}

	// 2. 更新状态
	updates := map[string]interface{}{
		"enabled":    enabled,
		"updated_at": time.Now(),
	}
	err = u.agentRepo.Update(ctx, agentModel.ID, updates)
	if err != nil {
		u.log.Errorf("agentUsecase.ToggleStatus.Update error: %v", err)
		return nil, code.ErrAgentUnKnown
	}

	// 3. 重新获取更新后的数据
	agentModel, err = u.agentRepo.GetByAgentID(ctx, agentID)
	if err != nil {
		u.log.Errorf("agentUsecase.ToggleStatus.GetByAgentID after update error: %v", err)
		return nil, code.ErrAgentUnKnown
	}

	return convert.AgentPo2Do(agentModel), nil
}
