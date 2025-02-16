package dao

import (
	"context"
	"webook/internal/domain"
	"webook/pkg/loggerx"

	"gorm.io/gorm"
)

type DraftArticleDao interface {
	Insert(ctx context.Context, article domain.Article) (int64, error)
}

type draftArticleDao struct {
	l  loggerx.Logger
	db *gorm.DB
}

func NewDraftArticleDao(l loggerx.Logger, db *gorm.DB) DraftArticleDao {
	return &draftArticleDao{
		l:  l,
		db: db,
	}
}

func (d *draftArticleDao) Insert(ctx context.Context, article domain.Article) (int64, error) {
	draftArticle := mapDomainToDbData(article)
	if err := d.db.WithContext(ctx).Create(&draftArticle).Error; err != nil {
		d.l.Error(ctx, "failed to insert draft article", loggerx.String("error", err.Error()))
		return 0, err
	}
	return draftArticle.Id, nil
}

type DraftArticle struct {
	Id      int64  `gorm:"column:id;primaryKey;autoIncrement"`
	UserId  int64  `gorm:"column:user_id"`
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func mapDomainToDbData(article domain.Article) DraftArticle {
	return DraftArticle{
		UserId:  article.UserId,
		Title:   article.Title,
		Content: article.Content,
	}
}
