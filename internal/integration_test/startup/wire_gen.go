// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package startup

import (
	"github.com/gin-gonic/gin"
	"webook/internal/interface/web"
	"webook/internal/repository"
	"webook/internal/repository/dao"
	"webook/internal/service"
	"webook/ioc"
)

// Injectors from wire.go:

func InitApiServer() *gin.Engine {
	db := ioc.InitDB()
	userDAO := dao.NewUserDAO(db)
	userRepository := repository.NewUserRepository(userDAO)
	userService := service.NewUserService(userRepository)
	userHandler := web.NewUserHandler(userService)
	engine := ioc.InitWebServer(userHandler)
	return engine
}
