package agent

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/entity"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/repo"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/convert"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/data/db/model"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/pkg/errors/code"
	"github.com/johnsonoklii/agentgo/pkg/utils"
)

type agentVersionUsecase struct {
	agentRepo        repo.AgentRepo
	agentVersionRepo repo.AgentVersionRepo
	log              *log.Helper
}

func NewAgentVersionUsecase(
	agentRepo repo.AgentRepo,
	agentVersionRepo repo.AgentVersionRepo,
	logger log.Logger,
) AgentVersionUsecase {
	return &agentVersionUsecase{
		agentRepo:        agentRepo,
		agentVersionRepo: agentVersionRepo,
		log:              log.NewHelper(logger),
	}
}

func (u *agentVersionUsecase) PublishAgent(ctx context.Context, req *PublishAgentRequest) (*entity.AgentVersion, error) {
	// 1. 获取当前 Agent 信息
	agentModel, err := u.agentRepo.GetByAgentID(ctx, req.AgentID)
	if err != nil {
		u.log.Errorf("PublishAgent.GetByAgentID error: %v", err)
		return nil, code.ErrAgentUnKnown
	}
	if agentModel == nil {
		return nil, code.ErrAgentNotFound
	}

	// 2. 检查版本号是否已存在
	existingVersion, err := u.agentVersionRepo.GetByAgentIDAndVersion(ctx, req.AgentID, req.VersionNumber)
	if err != nil {
		u.log.Errorf("PublishAgent.GetByAgentIDAndVersion error: %v", err)
		return nil, code.ErrAgentUnKnown
	}
	if existingVersion != nil {
		return nil, code.ErrAgentVersionExists
	}

	// 3. 创建新版本
	now := time.Now()
	versionModel := &model.AgentVersion{
		AgentVersionID:   utils.NewUUID(),
		AgentID:          req.AgentID,
		Name:             agentModel.Name,
		Avatar:           agentModel.Avatar,
		Description:      agentModel.Description,
		VersionNumber:    req.VersionNumber,
		SystemPrompt:     agentModel.SystemPrompt,
		WelcomeMessage:   agentModel.WelcomeMessage,
		ToolIds:          agentModel.ToolIds,
		KnowledgeBaseIds: agentModel.KnowledgeBaseIds,
		ChangeLog:        req.ChangeLog,
		PublishStatus:    model.PublishStatusReviewing, // 默认为审核中
		UserID:           req.UserID,
		ToolPresetParams: agentModel.ToolPresetParams,
		MultiModal:       agentModel.MultiModel,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	err = u.agentVersionRepo.Create(ctx, versionModel)
	if err != nil {
		u.log.Errorf("PublishAgent.Create error: %v", err)
		return nil, code.ErrAgentUnKnown
	}

	// 4. 更新 Agent 的 published_version 字段
	updates := map[string]interface{}{
		"published_version": req.VersionNumber,
	}
	err = u.agentRepo.Update(ctx, agentModel.ID, updates)
	if err != nil {
		u.log.Warnf("PublishAgent.UpdateAgent publishedVersion error: %v", err)
		// 不影响主流程，只记录日志
	}

	return convert.AgentVersionPo2Do(versionModel), nil
}

func (u *agentVersionUsecase) GetAgentVersions(ctx context.Context, agentID string) ([]*entity.AgentVersion, error) {
	versionModels, err := u.agentVersionRepo.ListByAgentID(ctx, agentID)
	if err != nil {
		u.log.Errorf("GetAgentVersions.ListByAgentID error: %v", err)
		return nil, code.ErrAgentUnKnown
	}

	versions := make([]*entity.AgentVersion, 0, len(versionModels))
	for _, versionModel := range versionModels {
		versions = append(versions, convert.AgentVersionPo2Do(versionModel))
	}

	return versions, nil
}

func (u *agentVersionUsecase) GetAgentVersion(ctx context.Context, agentID, versionNumber string) (*entity.AgentVersion, error) {
	versionModel, err := u.agentVersionRepo.GetByAgentIDAndVersion(ctx, agentID, versionNumber)
	if err != nil {
		u.log.Errorf("GetAgentVersion.GetByAgentIDAndVersion error: %v", err)
		return nil, code.ErrAgentUnKnown
	}

	if versionModel == nil {
		return nil, code.ErrAgentVersionNotFound
	}

	return convert.AgentVersionPo2Do(versionModel), nil
}

func (u *agentVersionUsecase) GetAgentLatestVersion(ctx context.Context, agentID string) (*entity.AgentVersion, error) {
	// 获取所有版本
	versionModels, err := u.agentVersionRepo.ListByAgentID(ctx, agentID)
	if err != nil {
		u.log.Errorf("GetAgentLatestVersion.ListByAgentID error: %v", err)
		return nil, code.ErrAgentUnKnown
	}

	if len(versionModels) == 0 {
		return nil, code.ErrAgentVersionNotFound
	}

	// 返回第一个（数据库已按创建时间倒序排序）
	return convert.AgentVersionPo2Do(versionModels[0]), nil
}
