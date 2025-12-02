package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/usecase/agent"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/convert"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/pkg/errors/code"
	"github.com/johnsonoklii/agentgo/pkg/jwt"

	pb "github.com/johnsonoklii/agentgo/apps/agentService/api/agent/v1"
)

type AgentService struct {
	pb.UnimplementedAgentServer
	agentUc        agent.AgentUsecase
	agentVersionUc agent.AgentVersionUsecase
	log            *log.Helper
}

func NewAgentService(agentUc agent.AgentUsecase, agentVersionUc agent.AgentVersionUsecase, logger log.Logger) *AgentService {
	return &AgentService{
		agentUc:        agentUc,
		agentVersionUc: agentVersionUc,
		log:            log.NewHelper(logger),
	}
}

func (s *AgentService) CreateAgent(ctx context.Context, req *pb.CreateAgentRequest) (*pb.CreateAgentResponse, error) {
	uid, ok := jwt.GetUserID(ctx)
	if !ok {
		return nil, code.ErrAgentNoAuth
	}

	agentEntity, err := s.agentUc.Create(ctx, uid, &agent.CreateAgentRequest{
		Name:         req.Name,
		Avatar:       req.Avatar,
		Description:  req.Description,
		SystemPrompt: req.SystemPrompt,
		UID:          uid,
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateAgentResponse{
		Code:    0,
		Message: "success",
		Data:    convert.AgentDo2To(agentEntity),
	}, nil
}

func (s *AgentService) GetAgentList(ctx context.Context, req *pb.GetAgentListRequest) (*pb.GetAgentListResponse, error) {
	agents, err := s.agentUc.GetAgentList(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	agentInfos := make([]*pb.AgentInfo, 0, len(agents))
	for _, agent := range agents {
		agentInfos = append(agentInfos, convert.AgentDo2To(agent))
	}

	return &pb.GetAgentListResponse{
		Code:    0,
		Message: "success",
		Data:    agentInfos,
	}, nil
}

func (s *AgentService) GetAgent(ctx context.Context, req *pb.GetAgentRequest) (*pb.GetAgentResponse, error) {
	agent, err := s.agentUc.GetAgent(ctx, req.AgentId)
	if err != nil {
		return nil, err
	}

	return &pb.GetAgentResponse{
		Code:    0,
		Message: "success",
		Data:    convert.AgentDo2To(agent),
	}, nil
}

func (s *AgentService) UpdateAgent(ctx context.Context, req *pb.UpdateAgentRequest) (*pb.UpdateAgentResponse, error) {
	uid, ok := jwt.GetUserID(ctx)
	if !ok {
		return nil, code.ErrAgentNoAuth
	}

	err := s.agentUc.Update(ctx, &agent.UpdateUserRequest{
		AgentId: req.AgentId,
		// Add other fields from request as needed
	})

	if err != nil {
		return nil, err
	}

	// Get updated modal
	updatedAgent, err := s.agentUc.GetAgent(ctx, req.AgentId)
	if err != nil {
		return nil, err
	}

	_ = uid // Use uid if needed for authorization check

	return &pb.UpdateAgentResponse{
		Code:    0,
		Message: "success",
		Data:    convert.AgentDo2To(updatedAgent),
	}, nil
}

func (s *AgentService) DeleteAgent(ctx context.Context, req *pb.DeleteAgentRequest) (*pb.DeleteAgentResponse, error) {
	uid, ok := jwt.GetUserID(ctx)
	if !ok {
		return nil, code.ErrAgentNoAuth
	}

	err := s.agentUc.Delete(ctx, req.AgentId)
	if err != nil {
		return nil, err
	}

	_ = uid // Use uid if needed for authorization check

	return &pb.DeleteAgentResponse{
		Code:    0,
		Message: "success",
	}, nil
}

func (s *AgentService) PublishAgent(ctx context.Context, req *pb.PublishAgentRequest) (*pb.PublishAgentResponse, error) {
	uid, ok := jwt.GetUserID(ctx)
	if !ok {
		return nil, code.ErrAgentNoAuth
	}

	version, err := s.agentVersionUc.PublishAgent(ctx, &agent.PublishAgentRequest{
		AgentID:       req.AgentId,
		VersionNumber: req.VersionNumber,
		ChangeLog:     req.ChangeLog,
		UserID:        uid,
	})

	if err != nil {
		return nil, err
	}

	return &pb.PublishAgentResponse{
		Code:    0,
		Message: "success",
		Data:    convert.AgentVersionDo2To(version),
	}, nil
}

func (s *AgentService) GetAgentVersions(ctx context.Context, req *pb.GetAgentVersionsRequest) (*pb.GetAgentVersionsResponse, error) {
	versions, err := s.agentVersionUc.GetAgentVersions(ctx, req.AgentId)
	if err != nil {
		return nil, err
	}

	versionInfos := make([]*pb.AgentVersionInfo, 0, len(versions))
	for _, version := range versions {
		versionInfos = append(versionInfos, convert.AgentVersionDo2To(version))
	}

	return &pb.GetAgentVersionsResponse{
		Code:    0,
		Message: "success",
		Data:    versionInfos,
	}, nil
}

func (s *AgentService) GetAgentVersion(ctx context.Context, req *pb.GetAgentVersionRequest) (*pb.GetAgentVersionResponse, error) {
	version, err := s.agentVersionUc.GetAgentVersion(ctx, req.AgentId, req.VersionNumber)
	if err != nil {
		return nil, err
	}

	return &pb.GetAgentVersionResponse{
		Code:    0,
		Message: "success",
		Data:    convert.AgentVersionDo2To(version),
	}, nil
}

func (s *AgentService) GetAgentLatestVersion(ctx context.Context, req *pb.GetAgentLatestVersionRequest) (*pb.GetAgentLatestVersionResponse, error) {
	version, err := s.agentVersionUc.GetAgentLatestVersion(ctx, req.AgentId)
	if err != nil {
		return nil, err
	}

	return &pb.GetAgentLatestVersionResponse{
		Code:    0,
		Message: "success",
		Data:    convert.AgentVersionDo2To(version),
	}, nil
}

func (s *AgentService) ToggleAgentStatus(ctx context.Context, req *pb.ToggleAgentStatusRequest) (*pb.ToggleAgentStatusResponse, error) {
	agent, err := s.agentUc.ToggleStatus(ctx, req.AgentId, req.Enabled)
	if err != nil {
		return nil, err
	}

	return &pb.ToggleAgentStatusResponse{
		Code:    0,
		Message: "success",
		Data:    convert.AgentDo2To(agent),
	}, nil
}
