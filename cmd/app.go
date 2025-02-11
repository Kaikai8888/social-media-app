package main

import "github.com/gin-gonic/gin"

type App struct {
	webServer *gin.Engine
}

func (a *App) Start() {
	a.webServer.Run(":8080")
}
