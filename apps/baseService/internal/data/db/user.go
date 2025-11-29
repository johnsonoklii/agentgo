package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/biz/repo"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/data/db/model"
	"gorm.io/gorm"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) repo.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Create(ctx context.Context, user *model.User) error {
	return r.data.DB.WithContext(ctx).Create(user).Error
}

func (r *userRepo) Delete(ctx context.Context, uid string) error {
	return r.data.DB.WithContext(ctx).Delete(&model.User{}, "uid = ?", uid).Error
}

func (r *userRepo) GetByMobile(ctx context.Context, mobile string) (*model.User, error) {
	user := &model.User{}
	err := r.data.DB.WithContext(ctx).Where("mobile = ?", mobile).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepo) Get(ctx context.Context, uid string) (*model.User, error) {
	user := &model.User{}
	err := r.data.DB.WithContext(ctx).Where("uid = ?", uid).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) Update(ctx context.Context, uid string, updates map[string]any) error {
	if _, ok := updates["updated_at"]; !ok {
		updates["updated_at"] = time.Now().UnixMilli()
	}

	return r.data.DB.WithContext(ctx).Model(&model.User{}).Where("uid = ?", uid).Updates(updates).Error
}
func (r *userRepo) UpdatePassword(ctx context.Context, mobile string, password string) error {
	return nil
}

func (r *userRepo) CheckMobileExist(ctx context.Context, mobile string) (bool, error) {
	userModel, err := r.GetByMobile(ctx, mobile)
	if err != nil {
		return false, err
	}
	return userModel == nil, err
}
