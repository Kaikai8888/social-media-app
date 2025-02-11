package main

import (
	"time"
	"webook/internal/interface/web"
	"webook/internal/interface/web/middleware"
	"webook/internal/repository"
	"webook/internal/repository/dao"
	"webook/internal/service"
	"webook/pkg/ginx/ginx/middleware/ratelimit"

	"strings"

	"webook/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDB()
	server := initWebServer()
	hdl := initUserHandler(db)
	hdl.RegisterRoutes(server)

	server.Run(":8080")
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Config.DBConfig.DSN))
	if err != nil {
		panic(err)
	}

	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}

func initUserHandler(db *gorm.DB) *web.UserHandler {
	userDAO := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(userDAO)
	svc := service.NewUserService(repo)
	hdl := web.NewUserHandler(svc)
	return hdl
}

func initWebServer() *gin.Engine {
	server := gin.Default()

	// CORS
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowOriginFunc: func(origin string) bool {
			return strings.Contains(origin, "localhost")
		},
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"x-jwt-token"}, // 允許前端訪問後端response中的x-jwt-token header
	}))

	redisClient := redis.NewClient(&redis.Options{
		Addr: config.Config.RedisConfig.Addr,
	})

	server.Use(ratelimit.NewBuilder(redisClient,
		time.Second, 1).Build())

	loginMiddleware := &middleware.LoginMiddlewareBuilder{}
	// check login
	server.Use(loginMiddleware.CheckLogin())

	return server
}
