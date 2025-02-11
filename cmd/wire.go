//go:build wireinject
// +build wireinject

package main

import (
	"webook/internal/interface/web"
	"webook/internal/repository"
	"webook/internal/repository/dao"
	"webook/internal/service"
	"webook/ioc"

	"github.com/google/wire"
)

func InitApp() *App {
	wire.Build(
		ioc.InitDB,

		dao.NewUserDAO,

		repository.NewUserRepository,

		service.NewUserService,

		web.NewUserHandler,

		ioc.InitWebServer,

		wire.Struct(new(App), "*"),
	)

	return new(App)
}
