package dao

import (
	"context"

	"gorm.io/gorm"
)

// dao: data access object
type UserDAO struct {
	db *gorm.DB
}

func (ud *UserDAO) Insert(ctx context.Context, u User) error {
	return ud.db.WithContext(ctx).Create(&u).Error
}

type User struct {
	ID       int64  `gorm:"primaryKey,autoIncrement"`
	Email    string `gorm:"unique"`
	Password string

	CreatedAt int64 // 用int64 or UTC 0 的 time.Time, 避免時區轉換相關的問題 (e.g. server 和 db時區不同)
	UpdatedAt int64
}
