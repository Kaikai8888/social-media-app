package ginx

import "github.com/gin-gonic/gin"

func Wrap[T any](handler func(ctx *gin.Context) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := handler(ctx); err != nil {
			ctx.JSON(500, gin.H{"message": err.Error()})
		}
	}
}
