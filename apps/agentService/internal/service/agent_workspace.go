package service

import (
	"context"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/entity"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/usecase/modal"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/usecase/provider"

	"github.com/go-kratos/kratos/v2/log"
	agentpb "github.com/johnsonoklii/agentgo/apps/agentService/api/agent/v1"
	pb "github.com/johnsonoklii/agentgo/apps/agentService/api/agent_workspace/v1"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/usecase/agent_workspace"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/convert"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/pkg/errors/code"
	"github.com/johnsonoklii/agentgo/pkg/jwt"
)

type AgentWorkspaceService struct {
	pb.UnimplementedAgentWorkspaceServer
	workspaceUc agent_workspace.AgentWorkspaceUsecase
	modalUc     modal.ModalUsecase
	providerUc  provider.ProviderUsecase
	log         *log.Helper
}

func NewAgentWorkspaceService(workspaceUc agent_workspace.AgentWorkspaceUsecase, providerUc provider.ProviderUsecase, modalUc modal.ModalUsecase, logger log.Logger) *AgentWorkspaceService {
	return &AgentWorkspaceService{
		workspaceUc: workspaceUc,
		modalUc:     modalUc,
		providerUc:  providerUc,
		log:         log.NewHelper(logger),
	}
}

func (s *AgentWorkspaceService) GetWorkspaceAgents(ctx context.Context, req *pb.GetWorkspaceAgentsRequest) (*pb.GetWorkspaceAgentsResponse, error) {
	agents, err := s.workspaceUc.GetWorkspaceAgents(ctx, req.WorkspaceId)
	if err != nil {
		return nil, err
	}

	agentInfos := make([]*agentpb.AgentInfo, 0, len(agents))
	for _, agent := range agents {
		agentInfos = append(agentInfos, convert.AgentDo2To(agent))
	}

	return &pb.GetWorkspaceAgentsResponse{
		Code:    0,
		Message: "success",
		Data:    agentInfos,
	}, nil
}

func (s *AgentWorkspaceService) AddAgentToWorkspace(ctx context.Context, req *pb.AddAgentToWorkspaceRequest) (*pb.AddAgentToWorkspaceResponse, error) {
	uid, ok := jwt.GetUserID(ctx)
	if !ok {
		return nil, code.ErrWorkspaceNoAuth
	}

	err := s.workspaceUc.AddAgentToWorkspace(ctx, uid, req.WorkspaceId, req.AgentId)
	if err != nil {
		return nil, err
	}

	return &pb.AddAgentToWorkspaceResponse{
		Code:    0,
		Message: "success",
	}, nil
}

func (s *AgentWorkspaceService) RemoveAgentFromWorkspace(ctx context.Context, req *pb.RemoveAgentFromWorkspaceRequest) (*pb.RemoveAgentFromWorkspaceResponse, error) {
	uid, ok := jwt.GetUserID(ctx)
	if !ok {
		return nil, code.ErrWorkspaceNoAuth
	}

	err := s.workspaceUc.RemoveAgentFromWorkspace(ctx, uid, req.WorkspaceId, req.AgentId)
	if err != nil {
		return nil, err
	}

	return &pb.RemoveAgentFromWorkspaceResponse{
		Code:    0,
		Message: "success",
	}, nil
}

// SetAgentModelConfig 设置Agent的模型配置
func (s *AgentWorkspaceService) UpdateAgentModelConfig(ctx context.Context, req *pb.UpdateModalConfigRequest) (*pb.UpdateModalConfigResponse, error) {
	// 验证用户身份
	uid, ok := jwt.GetUserID(ctx)
	if !ok {
		return nil, code.ErrWorkspaceNoAuth
	}

	modalEntity, err := s.modalUc.GetModalByID(ctx, req.ModalId)
	if err != nil {
		return nil, err
	}

	if !modalEntity.IsActive() {
		return nil, code.ErrModalNoActive
	}

	providerEntity, err := s.providerUc.GetProviderByID(ctx, modalEntity.ProviderID)
	if err != nil {
		return nil, err
	}
	if !providerEntity.IsActive() {
		return nil, code.ErrProviderNoActive
	}

	err = s.workspaceUc.UpdateAgentModelConfig(ctx, &entity.AgentWorkspace{
		UID:     uid,
		AgentID: req.AgentId,
		LLMModalConfig: entity.LLMModalConfig{
			ModalId:                   req.ModalId,
			Temperature:               float64(req.Temperature),
			TopP:                      float64(req.TopP),
			MaxTokens:                 int(req.MaxTokens),
			ReserveRatio:              float64(req.ReserveRatio),
			SummaryThreshold:          int(req.SummaryThreshold),
			TokenOverflowStrategyEnum: entity.TokenOverflowStrategyEnum(req.StrategyType),
		},
	})

	if err != nil {
		return nil, code.ErrWorkspaceUnknown
	}

	return &pb.UpdateModalConfigResponse{
		Code:    0,
		Message: "success",
	}, nil
}
