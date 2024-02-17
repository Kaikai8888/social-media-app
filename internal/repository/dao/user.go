package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrDuplicateEmail = errors.New("duplicated email")
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

// dao: data access object
type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (ud *UserDAO) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.CreatedAt = now
	u.UpdatedAt = now
	err := ud.db.WithContext(ctx).Create(&u).Error
	if me, ok := err.(*mysql.MySQLError); ok {
		const duplicatedErr uint16 = 1062
		if me.Number == duplicatedErr {
			return ErrDuplicateEmail
		}
	}

	return err
}

func (ud *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := ud.db.WithContext(ctx).Where("email=?", email).First(&u).Error
	return u, err
}

type User struct {
	ID       int64  `gorm:"primaryKey,autoIncrement"`
	Email    string `gorm:"unique"`
	Password string

	CreatedAt int64 // 用int64 or UTC 0 的 time.Time, 避免時區轉換相關的問題 (e.g. server 和 db時區不同)
	UpdatedAt int64
}
