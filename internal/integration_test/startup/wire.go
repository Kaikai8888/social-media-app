//go:build wireinject
// +build wireinject

package startup

import (
	"webook/internal/interface/web"
	"webook/internal/repository"
	"webook/internal/repository/dao"
	"webook/internal/service"
	"webook/ioc"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitApiServer() *gin.Engine {
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
	)

	return new(gin.Engine)
}
