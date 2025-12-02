package convert

import (
	pb "github.com/johnsonoklii/agentgo/apps/baseService/api/user/v1"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/biz/entity"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/data/db/model"
)

func UserPo2Do(model *model.User, token string) *entity.User {
	if model == nil {
		return nil
	}

	return &entity.User{
		UID:       model.UID,
		Mobile:    model.Mobile,
		UserName:  model.UserName,
		Gender:    model.Gender,
		Deleted:   model.Deleted,
		Token:     token,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func UserDo2To(userDo *entity.User) *pb.UserInfo {
	if userDo == nil {
		return nil
	}
	return &pb.UserInfo{
		UID:       userDo.UID,
		Mobile:    userDo.Mobile,
		UserName:  userDo.UserName,
		Gender:    int32(userDo.Gender),
		CreatedAt: userDo.CreatedAt.UnixMilli(),
	}
}
