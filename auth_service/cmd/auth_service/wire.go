// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"auth_service/internal/biz"
	"auth_service/internal/conf"
	"auth_service/internal/data"
	"auth_service/internal/server"
	"auth_service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
