package convert

import (
	pb "github.com/johnsonoklii/agentgo/apps/agentService/api/agent/v1"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/entity"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/data/db/model"
)

// AgentVersionPo2Do Model 转 Entity
func AgentVersionPo2Do(m *model.AgentVersion) *entity.AgentVersion {
	if m == nil {
		return nil
	}

	return &entity.AgentVersion{
		AgentVersionID:   m.AgentVersionID,
		AgentID:          m.AgentID,
		Name:             m.Name,
		Avatar:           m.Avatar,
		Description:      m.Description,
		VersionNumber:    m.VersionNumber,
		SystemPrompt:     m.SystemPrompt,
		WelcomeMessage:   m.WelcomeMessage,
		ToolIds:          m.ToolIds,
		KnowledgeBaseIds: m.KnowledgeBaseIds,
		ChangeLog:        m.ChangeLog,
		PublishStatus:    m.PublishStatus,
		RejectReason:     m.RejectReason,
		ReviewTime:       m.ReviewTime,
		PublishedAt:      m.PublishedAt,
		UserID:           m.UserID,
		ToolPresetParams: m.ToolPresetParams,
		MultiModal:       m.MultiModal,
		CreatedAt:        m.CreatedAt,
		UpdatedAt:        m.UpdatedAt,
	}
}

// AgentVersionDo2To Entity 转 Protobuf
func AgentVersionDo2To(e *entity.AgentVersion) *pb.AgentVersionInfo {
	if e == nil {
		return nil
	}

	var reviewTime int64
	if e.ReviewTime != nil {
		reviewTime = e.ReviewTime.UnixMilli()
	}

	var publishedAt int64
	if e.PublishedAt != nil {
		publishedAt = e.PublishedAt.UnixMilli()
	}

	return &pb.AgentVersionInfo{
		AgentVersionID:   e.AgentVersionID,
		AgentId:          e.AgentID,
		Name:             e.Name,
		Avatar:           e.Avatar,
		Description:      e.Description,
		VersionNumber:    e.VersionNumber,
		SystemPrompt:     e.SystemPrompt,
		WelcomeMessage:   e.WelcomeMessage,
		ToolIds:          e.ToolIds,
		KnowledgeBaseIds: e.KnowledgeBaseIds,
		ChangeLog:        e.ChangeLog,
		PublishStatus:    int32(e.PublishStatus),
		RejectReason:     e.RejectReason,
		ReviewTime:       reviewTime,
		PublishedAt:      publishedAt,
		UserId:           e.UserID,
		// ToolPresetParams: e.ToolPresetParams, // TODO: 需要转换
		MultiModal: e.MultiModal,
		CreatedAt:  e.CreatedAt.UnixMilli(),
		UpdatedAt:  e.UpdatedAt.UnixMilli(),
	}
}
