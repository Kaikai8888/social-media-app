package repository

import (
	"context"
	"social-media-app/internal/domain"
	"social-media-app/internal/repository/dao"
	"social-media-app/pkg/loggerx"
)

type DraftArticleRepository interface {
	Create(ctx context.Context, article domain.Article) (int64, error)
}

type draftArticleRepository struct {
	l               loggerx.Logger
	draftArticleDao dao.DraftArticleDao
}

func NewDraftArticleRepository(l loggerx.Logger, dao dao.DraftArticleDao) DraftArticleRepository {
	return &draftArticleRepository{
		l:               l,
		draftArticleDao: dao,
	}
}

func (r *draftArticleRepository) Create(ctx context.Context, article domain.Article) (int64, error) {
	return r.draftArticleDao.Insert(ctx, article)
}
