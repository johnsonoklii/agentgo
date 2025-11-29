package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/johnsonoklii/agentgo/apps/baseService/api/auth/v1"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/biz/usecase/auth"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/pkg/errorx/code"
	"github.com/johnsonoklii/agentgo/pkg/jwt"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	authUc auth.AuthUsecase
	log    *log.Helper
}

func NewAuthService(authUc auth.AuthUsecase, logger log.Logger) *AuthService {
	return &AuthService{
		authUc: authUc,
		log:    log.NewHelper(logger),
	}
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if !isValidMobile(req.Mobile) {
		return nil, code.ErrUserInfo
	}

	userEntity, err := s.authUc.Login(ctx, req.Mobile, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Token:   userEntity.Token,
		Message: "success",
		Code:    0,
	}, nil
}

func (s *AuthService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	uid, ok := jwt.GetUserID(ctx)
	if !ok {
		return nil, code.ErrUserNoAuth
	}

	fmt.Println()
	err := s.authUc.Logout(ctx, uid)
	if err != nil {
		return nil, err
	}

	return &pb.LogoutResponse{
		Message: "success",
		Code:    0,
	}, nil
}
