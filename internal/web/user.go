package web

import "github.com/gin-gonic/gin"

type UserHandler struct {
}

func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.POST("/users/signup", h.SignUp)
	ug.POST("/users/login", h.Login)
	ug.POST("/users/edit", h.Edit)
	ug.GET("/users/profile", h.Profile)
}

func (h *UserHandler) SignUp(ctx *gin.Context) {

}

func (h *UserHandler) Login(ctx *gin.Context) {

}

func (h *UserHandler) Edit(ctx *gin.Context) {

}

func (h *UserHandler) Profile(ctx *gin.Context) {

}
