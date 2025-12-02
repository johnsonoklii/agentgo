package user

import (
	"context"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/biz/entity"
)

type RegisterRequest struct {
	UserName string
	Password string
	Mobile   string
	Gender   int8
}

type UpdateUserRequest struct {
	UID      string
	UserName string
	Password string
	Mobile   string
	Gender   int8
}

type UserUsecase interface {
	Register(ctx context.Context, req *RegisterRequest) error
	Delete(ctx context.Context, uid string) error
	UnDelete(ctx context.Context, uid string) error
	GetUserByMobile(ctx context.Context, mobile string) (*entity.User, error)
	GetUserByID(ctx context.Context, uid string) (*entity.User, error)
	GetUsersByIDs(ctx context.Context, uids []string) ([]*entity.User, error)
	Update(ctx context.Context, req *UpdateUserRequest) error
	UpdatePassword(ctx context.Context, mobile string, password string) error
}
