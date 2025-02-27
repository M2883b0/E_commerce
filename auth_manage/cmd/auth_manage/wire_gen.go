// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"auth_manage/internal/biz"
	"auth_manage/internal/conf"
	"auth_manage/internal/data"
	"auth_manage/internal/server"
	"auth_manage/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	authRepo := data.NewAuthRepo(dataData, logger)
	authUsecase := biz.NewAuthUsecase(authRepo, logger)
	authService := service.NewAuthService(authUsecase)
	grpcServer := server.NewGRPCServer(confServer, authService, logger)
	app := newApp(confData, logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
