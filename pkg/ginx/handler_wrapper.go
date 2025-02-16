package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

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
			ctx.JSON(401, gin.H{"code": "401", "message": "unauthorized"})
			return
		}

		var req T
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(400, map[string]any{"code": "400", "message": "failed to parse body"})
			return
		}

		validate := validator.New()
		if err := validate.Struct(req); err != nil {
			ctx.JSON(400, map[string]any{"code": "400", "message": "invalid request"})
			return
		}

		handler(ctx, req, u.(U))
	}
}
