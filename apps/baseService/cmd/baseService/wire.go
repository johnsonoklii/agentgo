//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/biz"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/conf"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/data/db"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/server"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, db.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
