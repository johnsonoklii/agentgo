package repo

import (
	"context"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/data/db/model"
)

type UserRepo interface {
	Create(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, uid string) error
	GetByMobile(ctx context.Context, mobile string) (*model.User, error)
	Get(ctx context.Context, uid string) (*model.User, error)
	Update(ctx context.Context, uid string, updates map[string]any) error
	UpdatePassword(ctx context.Context, mobile string, password string) error
	CheckMobileExist(ctx context.Context, mobile string) (bool, error)
}
