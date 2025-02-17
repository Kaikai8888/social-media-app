package ioc

import (
	"strings"
	"time"
	"social-media-app/internal/interface/web"
	"social-media-app/internal/interface/web/middleware"
	"social-media-app/internal/repository"
	"social-media-app/internal/repository/dao"
	"social-media-app/internal/service"
	"social-media-app/pkg/ginx/middleware/ratelimit"
	"social-media-app/pkg/ginx/middleware/trace_id_allocator"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func InitUserHandler(db *gorm.DB) *web.UserHandler {
	userDAO := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(userDAO)
	svc := service.NewUserService(repo)
	hdl := web.NewUserHandler(svc)
	return hdl
}

func InitWebServer(redisClient redis.Cmdable, userHandler *web.UserHandler, articleHandler *web.ArticleHandler) *gin.Engine {
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

	server.Use(ratelimit.NewBuilder(redisClient,
		time.Second, 1).Build())

	// middlewares
	loginMiddleware := &middleware.LoginMiddlewareBuilder{}
	server.Use(loginMiddleware.CheckLogin())
	traceIdAllocator := &trace_id_allocator.Builder{}
	server.Use(traceIdAllocator.Build())

	// register routes
	userHandler.RegisterRoutes(server)
	articleHandler.RegisterRoutes(server)

	return server
}
