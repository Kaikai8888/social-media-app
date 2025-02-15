package ginx

import "github.com/gin-gonic/gin"

type MiddlewareBuilder interface {
	Build() gin.HandlerFunc
	IgnorePaths(paths ...string) MiddlewareBuilder
}
