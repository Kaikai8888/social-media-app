//go:build wireinject
// +build wireinject

package startup

import (
	"social-media-app/internal/interface/web"
	"social-media-app/internal/repository"
	"social-media-app/internal/repository/dao"
	"social-media-app/internal/service"
	"social-media-app/ioc"

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
