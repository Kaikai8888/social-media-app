package ginx

import "github.com/gin-gonic/gin"

const FieldUser = "user"

func Wrap[U any](handler func(ctx *gin.Context, userClaims U)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u, ok := ctx.Get(FieldUser)
		if !ok {
			ctx.JSON(401, gin.H{"message": "unauthorized"})
		}

		handler(ctx, u.(U))
	}
}

func WrapRequest[T, U any](handler func(ctx *gin.Context, request T, userClaims U)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u, ok := ctx.Get(FieldUser)
		if !ok {
			ctx.JSON(401, gin.H{"message": "unauthorized"})
		}

		var req T
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(400, gin.H{"message": "failed to parse body"})
			return
		}

		handler(ctx, req, u.(U))
	}
}
