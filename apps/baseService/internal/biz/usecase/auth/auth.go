package auth

import (
	"context"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/biz/entity"
)

type AuthUsecase interface {
	Login(ctx context.Context, mobile string, password string) (*entity.User, error)
	Logout(ctx context.Context, uid string) error
}
