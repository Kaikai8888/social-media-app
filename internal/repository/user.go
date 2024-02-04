package repository

import (
	"context"
	"webook/internal/domain"
	"webook/internal/repository/dao"
)

type UserRepository struct {
	dao *dao.UserDAO
}

func (ur *UserRepository) Create(ctx context.Context, user domain.User) error {
	return ur.dao.Insert(ctx, dao.User{
		Email:    user.Email,
		Password: user.Password,
	})
}
