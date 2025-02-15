package trace

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const FieldTraceId = "trace_id"

func NewTraceId() string {
	return uuid.New().String()
}

func GetTraceId(ctx context.Context) string {
	switch c := ctx.(type) {
	case *gin.Context:
		return c.GetString(FieldTraceId)
	default:
		return ctx.Value(FieldTraceId).(string)
	}
}
