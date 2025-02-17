package service

import (
	"context"
	"social-media-app/internal/domain"
	"social-media-app/internal/repository"
	"social-media-app/pkg/loggerx"
)

type ArticleService interface {
	Create(ctx context.Context, article domain.Article) (int64, error)
}

type articleService struct {
	l    loggerx.Logger
	repo repository.DraftArticleRepository
}

func NewArticleService(l loggerx.Logger, repo repository.DraftArticleRepository) ArticleService {
	return &articleService{
		l:    l,
		repo: repo,
	}
}

func (s *articleService) Create(ctx context.Context, article domain.Article) (int64, error) {
	return s.repo.Create(ctx, article)
}
