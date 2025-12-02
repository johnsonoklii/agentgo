package service

import (
	"context"

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
	log         *log.Helper
}

func NewAgentWorkspaceService(workspaceUc agent_workspace.AgentWorkspaceUsecase, logger log.Logger) *AgentWorkspaceService {
	return &AgentWorkspaceService{
		workspaceUc: workspaceUc,
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
		return nil, code.ErrAgentNoAuth
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
		return nil, code.ErrAgentNoAuth
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
