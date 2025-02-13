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

		dao.NewUserDAO,

		repository.NewUserRepository,

		service.NewUserService,

		web.NewUserHandler,

		ioc.InitWebServer,
	)

	return new(gin.Engine)
}
