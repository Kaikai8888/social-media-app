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
		ioc.InitLogger,

		dao.NewUserDAO,
		dao.NewDraftArticleDao,

		repository.NewUserRepository,
		repository.NewDraftArticleRepository,

		service.NewUserService,
		service.NewArticleService,

		web.NewUserHandler,
		web.NewArticleHandler,

		ioc.InitWebServer,

		wire.Struct(new(App), "*"),
	)

	return new(App)
}
