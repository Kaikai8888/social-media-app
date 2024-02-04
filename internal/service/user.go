package service

import (
	"context"
	"webook/internal/domain"
	"webook/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func (u *UserService) Signup(ctx context.Context, user domain.User) {

}
