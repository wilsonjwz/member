package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"user/internal/data/ent"
)

type User struct {
	Hello string
}

type UserRepo interface {
	GetUserInfo(context.Context, string) (*ent.User, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) FindByUserName(ctx context.Context, username string) (*ent.User, error) {
	return uc.repo.GetUserInfo(ctx, username)
}
