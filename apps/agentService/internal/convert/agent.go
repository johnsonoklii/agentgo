package convert

import (
	pb "github.com/johnsonoklii/agentgo/apps/agentService/api/agent/v1"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/entity"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/data/db/model"
)

func AgentPo2Do(model *model.Agent) *entity.Agent {
	if model == nil {
		return nil
	}

	return &entity.Agent{
		AgentID:          model.AgentID,
		Name:             model.Name,
		Avatar:           model.Avatar,
		Description:      model.Description,
		SystemPrompt:     model.SystemPrompt,
		ToolIds:          model.ToolIds,
		KnowledgeBaseIds: model.KnowledgeBaseIds,
		PublishedVersion: model.PublishedVersion,
		Enabled:          model.Enabled,
		UID:              model.UID,
		ToolPresetParams: model.ToolPresetParams,
		MultiModel:       model.MultiModel,
		CreatedAt:        model.CreatedAt,
		UpdatedAt:        model.UpdatedAt,
	}
}

func AgentDo2To(userDo *entity.Agent) *pb.AgentInfo {
	if userDo == nil {
		return nil
	}
	return &pb.AgentInfo{
		AgentID:          userDo.AgentID,
		Name:             userDo.Name,
		Avatar:           userDo.Avatar,
		Description:      userDo.Description,
		SystemPrompt:     userDo.SystemPrompt,
		ToolIds:          userDo.ToolIds,
		KnowledgeBaseIds: userDo.KnowledgeBaseIds,
		PublishedVersion: userDo.PublishedVersion,
		Enabled:          userDo.Enabled,
		Uid:              userDo.UID,
		//ToolPresetParams: userDo.ToolPresetParams,
		MultiModel: userDo.MultiModel,
		CreatedAt:  userDo.CreatedAt.UnixMilli(),
		UpdatedAt:  userDo.UpdatedAt.UnixMilli(),
	}
}
