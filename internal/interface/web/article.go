package web

import (
	"social-media-app/internal/domain"
	"social-media-app/internal/service"
	"social-media-app/pkg/ginx"
	"social-media-app/pkg/loggerx"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	l          loggerx.Logger
	articleSvc service.ArticleService
}

func NewArticleHandler(l loggerx.Logger, articleSvc service.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		l:          l,
		articleSvc: articleSvc,
	}
}

func (h *ArticleHandler) RegisterRoutes(server *gin.Engine) {
	ag := server.Group("/articles")
	ag.POST("", ginx.WrapRequest[CreateUserReq, UserClaims](h.Create))
}

func (h *ArticleHandler) Create(ctx *gin.Context, req CreateUserReq, uc UserClaims) {
	article := domain.Article{
		UserId:  uc.Uid,
		Title:   req.Title,
		Content: req.Content,
	}
	id, err := h.articleSvc.Create(ctx, article)
	if err != nil {
		h.l.Error(ctx, "failed to create article", loggerx.Error(err))
		ctx.JSON(500, Response{Message: "failed to create article", Code: "500"})
		return
	}

	type Data struct {
		Id int64 `json:"id"`
	}

	ctx.JSON(200, Response{Message: "success", Code: "200", Data: Data{Id: id}})
}

type CreateUserReq struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}
