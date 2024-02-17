package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginMiddlewareBuilder struct{}

func (m *LoginMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/users/signup" || c.Request.URL.Path == "/users/login" {
			return
		}

		sess := sessions.Default(c)
		if sess.Get("userId") == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}

}
