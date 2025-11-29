package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/biz/entity"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/biz/repo"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/convert"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/data/db/model"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/pkg/encrypt"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/pkg/errorx/code"

	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type userUsecase struct {
	userRepo repo.UserRepo
	log      *log.Helper
}

func NewUserUsecase(userRepo repo.UserRepo, logger log.Logger) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
		log:      log.NewHelper(logger),
	}
}

func (uc *userUsecase) Register(ctx context.Context, req *RegisterRequest) error {
	userModel, err := uc.GetUserByMobile(ctx, req.Mobile)
	if err != nil {
		uc.log.Errorf("userUsecase.Register.GetUserByMobile error: %v", err)
		return code.ErrUserUnKnown
	}

	hashedPassword, err := encrypt.Hash(req.Password)
	if err != nil {
		uc.log.Errorf("userUsecase.Register.hashPassword error: %v", err)
		return code.ErrUserUnKnown
	}

	userName := req.UserName
	if userName == "" {
		userName = req.Mobile
	}

	now := time.Now().UnixMilli()

	if userModel == nil {
		// 创建
		err = uc.userRepo.Create(ctx, &model.User{
			UID:       uuid.NewString(),
			Mobile:    req.Mobile,
			UserName:  req.UserName,
			Password:  hashedPassword,
			Gender:    req.Gender,
			CreatedAt: now,
			UpdatedAt: now,
		})
		if err != nil {
			uc.log.Errorf("userUsecase.Register.Create error: %v", err)
			return code.ErrUserUnKnown
		}

	} else if userModel.Deleted {
		// 更新
		updates := map[string]any{
			"mobile":     req.Mobile,
			"user_name":  req.UserName,
			"password":   hashedPassword,
			"gender":     req.Gender,
			"deleted":    false,
			"created_at": now,
			"updated_at": now,
		}

		err = uc.userRepo.Update(ctx, userModel.UID, updates)
		if err != nil {
			uc.log.Errorf("userUsecase.Register.Update error: %v", err)
			return code.ErrUserUnKnown
		}
	} else {
		log.Errorf("userUsecase.Register error: %v already register", req.Mobile)
		return code.ErrUserExisted
	}

	return nil
}

func (uc *userUsecase) Delete(ctx context.Context, uid string) error {
	updates := map[string]any{
		"updated_at": time.Now().UnixMilli(),
	}
	updates["deleted"] = true

	err := uc.userRepo.Update(ctx, uid, updates)
	if err != nil {
		uc.log.Errorf("userUsecase.Delete.Update error: %v", err)
		return code.ErrUserUnKnown
	}

	return nil
}

func (uc *userUsecase) UnDelete(ctx context.Context, uid string) error {
	updates := map[string]any{
		"deleted": time.Now().UnixMilli(),
	}
	updates["deleted"] = false

	err := uc.userRepo.Update(ctx, uid, updates)
	if err != nil {
		uc.log.Errorf("userUsecase.UnDelete.Update error: %v", err)
		return err
	}

	return nil
}

func (uc *userUsecase) GetUserByMobile(ctx context.Context, mobile string) (*entity.User, error) {
	userModel, err := uc.userRepo.GetByMobile(ctx, mobile)
	if err != nil {
		uc.log.Errorf("userUsecase.GetUserByMobile.GetUserByMobile error: %v", err)
		return nil, err
	}

	return convert.UserPo2Do(userModel, ""), nil
}

func (uc *userUsecase) GetUserByID(ctx context.Context, uid string) (*entity.User, error) {
	return nil, nil
}

func (uc *userUsecase) GetUsersByIDs(ctx context.Context, uids []string) ([]*entity.User, error) {
	return nil, nil
}
func (uc *userUsecase) Update(ctx context.Context, req *UpdateUserRequest) error {
	return nil
}
func (uc *userUsecase) UpdatePassword(ctx context.Context, mobile string, password string) error {
	return uc.userRepo.UpdatePassword(ctx, mobile, password)
}
