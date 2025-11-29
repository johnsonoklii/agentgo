package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz/repo"

	pb "github.com/johnsonoklii/agentgo/apps/agentService/api/agent/v1"
)

type AgentService struct {
	pb.UnimplementedAgentServer
	agentRepo repo.AgentRepo
	log       *log.Helper
}

func NewAgentService(agentRepo repo.AgentRepo, logger log.Logger) *AgentService {
	return &AgentService{
		agentRepo: agentRepo,
		log:       log.NewHelper(logger),
	}
}

func (s *AgentService) CreateAgent(ctx context.Context, req *pb.CreateAgentRequest) (*pb.CreateAgentResponse, error) {

	return &pb.CreateAgentResponse{}, nil
}
func (s *AgentService) GetAgentList(ctx context.Context, req *pb.GetAgentListRequest) (*pb.GetAgentListResponse, error) {
	return &pb.GetAgentListResponse{}, nil
}
func (s *AgentService) GetAgent(ctx context.Context, req *pb.GetAgentRequest) (*pb.GetAgentResponse, error) {
	return &pb.GetAgentResponse{}, nil
}
func (s *AgentService) UpdateAgent(ctx context.Context, req *pb.UpdateAgentRequest) (*pb.UpdateAgentResponse, error) {
	return &pb.UpdateAgentResponse{}, nil
}
func (s *AgentService) DeleteAgent(ctx context.Context, req *pb.DeleteAgentRequest) (*pb.DeleteAgentResponse, error) {
	return &pb.DeleteAgentResponse{}, nil
}
