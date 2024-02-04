package main

import (
	"webook/internal/web"

	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	hdl := web.NewUserHandler()
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

	hdl.RegisterRoutes(server)

	server.Run(":8080")
}
