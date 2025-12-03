package auth

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/biz/entity"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/biz/repo"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/convert"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/data/db"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/pkg/errors/code"
	"github.com/johnsonoklii/agentgo/pkg/jwt"
	"github.com/johnsonoklii/agentgo/pkg/utils/encrypt"
)

type authUsecase struct {
	userRepo repo.UserRepo
	data     *db.Data
	log      *log.Helper
}

func NewAuthUsecase(userRepo repo.UserRepo, data *db.Data, logger log.Logger) AuthUsecase {
	return &authUsecase{
		userRepo: userRepo,
		data:     data,
		log:      log.NewHelper(logger),
	}
}

func (uc *authUsecase) Login(ctx context.Context, mobile string, password string) (*entity.User, error) {
	userModel, err := uc.userRepo.GetByMobile(ctx, mobile)
	if err != nil {
		uc.log.Errorf("userUsecase.Login.GetUserByMobile error: %v", err)
		return nil, code.ErrUserUnKnown
	}
	if userModel == nil {
		return nil, code.ErrUserNotFound
	}

	if userModel.Deleted {
		return nil, code.ErrUserNotFound
	}

	ok, err := encrypt.Verify(password, userModel.Password)
	if !ok {
		uc.log.Errorf("userUsecase.Login.VerifyPassword error: %v", err)
		return nil, code.ErrUserUnKnown
	}

	// jwt token
	token, err := jwt.GenerateToken(userModel.UID)
	if err != nil {
		uc.log.Errorf("userUsecase.Login.GenerateToken error: %v", err)
		return nil, code.ErrUserUnKnown
	}

	key := jwt.GetRDSKeyToken(userModel.UID)
	err = uc.data.BizRDB.Set(ctx, key, token, jwt.JwtExpire).Err()
	if err != nil {
		uc.log.Errorf("userUsecase.Login BizRDB.Set error: %v", err)
		return nil, code.ErrUserUnKnown
	}

	return convert.UserPo2Do(userModel, token), nil
}

func (uc *authUsecase) Logout(ctx context.Context, uid string) error {
	key := jwt.GetRDSKeyToken(uid)
	err := uc.data.BizRDB.Del(ctx, key).Err()
	if err != nil {
		uc.log.Errorf("userUsecase.Logout.Del error: %v", err)
		return code.ErrUserUnKnown
	}

	return nil
}
