// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"uuid_service/internal/biz"
	"uuid_service/internal/conf"
	"uuid_service/internal/data"
	"uuid_service/internal/server"
	"uuid_service/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	uuidRepo := data.NewUuidRepo(dataData, logger)
	uuidUsecase := biz.NewUuidUsecase(uuidRepo, logger)
	uuidService := service.NewUuidService(uuidUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, uuidService, logger)
	grpcServer := server.NewGRPCServer(confServer, uuidService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
