package trace_id_allocator

import (
	"webook/pkg/ginx"
	"webook/pkg/trace"

	"github.com/emirpasic/gods/v2/sets"
	"github.com/emirpasic/gods/v2/sets/hashset"
	"github.com/gin-gonic/gin"
)

const fieldTraceId = trace.FieldTraceId

type Builder struct {
	ignorePaths sets.Set[string]
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) IgnorePaths(paths ...string) ginx.MiddlewareBuilder {
	if b.ignorePaths == nil {
		b.ignorePaths = hashset.New[string]()
	}
	for _, path := range paths {
		b.ignorePaths.Add(path)
	}
	return b
}

func (b *Builder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if b.ignorePaths != nil && b.ignorePaths.Contains(ctx.Request.URL.Path) {
			ctx.Next()
			return
		}
		traceId := ctx.GetHeader("X-Request-Id")
		if traceId == "" {
			traceId = trace.NewTraceId()
		}
		ctx.Set(fieldTraceId, traceId)
		ctx.Next()
	}
}
