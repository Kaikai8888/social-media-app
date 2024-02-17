package main

import (
	"webook/internal/repository"
	"webook/internal/repository/dao"
	"webook/internal/service"
	"webook/internal/web"
	"webook/internal/web/middleware"

	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
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
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
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
		AllowHeaders:     []string{"Content-Type"},
	}))

	store, err := redis.NewStore(16, "tcp", "localhost:6379", "", []byte("k6CswdUm75WKcbM68UQUuxVsHSpTCwgK"), []byte("eF1`yQ9>yT1`tH1,sJ0.zD8;mZ9~nC6("))
	if err != nil {
		panic(err)
	}

	// get session
	server.Use(sessions.Sessions("ssid", store)) // name: cookie name

	loginMiddleware := &middleware.LoginMiddlewareBuilder{}
	// check login
	server.Use(loginMiddleware.CheckLogin())

	return server
}
