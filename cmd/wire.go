//go:build wireinject
// +build wireinject

package main

import (
	"social-media-app/internal/interface/web"
	"social-media-app/internal/repository"
	"social-media-app/internal/repository/dao"
	"social-media-app/internal/service"
	"social-media-app/ioc"

	"github.com/google/wire"
)

func InitApp() *App {
	wire.Build(
		ioc.InitDB,
		ioc.InitRedis,
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
