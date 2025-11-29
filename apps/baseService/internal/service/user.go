package service

import (
	"context"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/biz/usecase/user"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/pkg/errorx/code"
	"github.com/johnsonoklii/agentgo/pkg/jwt"
	"regexp"

	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/johnsonoklii/agentgo/apps/baseService/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
	userUc user.UserUsecase
	log    *log.Helper
}

func NewUserService(userUc user.UserUsecase, logger log.Logger) *UserService {
	return &UserService{
		userUc: userUc,
		log:    log.NewHelper(logger),
	}
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if !isValidMobile(req.Mobile) {
		return nil, code.ErrUserInfo
	}

	err := s.userUc.Register(ctx, &user.RegisterRequest{
		UserName: req.UserName,
		Mobile:   req.Mobile,
		Password: req.Password,
		Gender:   int8(req.Gender),
	})
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		Message: "success",
		Code:    0,
	}, nil
}

func (s *UserService) UnRegister(ctx context.Context, req *pb.UnRegisterRequest) (*pb.UnRegisterResponse, error) {
	UID, ok := jwt.GetUserID(ctx)
	if !ok {
		return nil, code.ErrUserUnKnown
	}

	if req.Uid != UID {
		return nil, code.ErrUserNoAuth
	}

	err := s.userUc.Delete(ctx, UID)
	if err != nil {
		return nil, err
	}
	return &pb.UnRegisterResponse{
		Message: "success",
		Code:    0,
	}, nil
}

func (s *UserService) Healthz(ctx context.Context, req *pb.HealthzRequest) (*pb.HealthzResponse, error) {
	return &pb.HealthzResponse{
		Code:    200,
		Message: "ok",
	}, nil
}

func isValidMobile(mobile string) bool {
	pattern := `^1[3-9]\d{9}$`
	match, _ := regexp.MatchString(pattern, mobile)
	return match
}
