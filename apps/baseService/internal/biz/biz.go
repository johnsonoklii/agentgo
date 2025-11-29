package biz

import (
	"github.com/google/wire"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/biz/usecase/auth"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/biz/usecase/user"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(user.NewUserUsecase, auth.NewAuthUsecase)
