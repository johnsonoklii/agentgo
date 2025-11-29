//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/biz"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/conf"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/data/repo"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/server"
	"github.com/johnsonoklii/agentgo/apps/agentService/internal/service"
	"github.com/johnsonoklii/agentgo/pkg/jwt"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
//func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
//	panic(wire.Build(server.ProviderSet, repo.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
//}

func wireApp(*conf.Server, *conf.Data, *conf.Registry, *jwt.Options, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, repo.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
