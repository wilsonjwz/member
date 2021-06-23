package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"user/internal/data/ent"
)

//User biz.User
type User struct {
	Hello string
}

//UserRepo 用户数据操作接口
type UserRepo interface {
	GetUserInfo(context.Context, string) (*ent.User, error)
}

//UserUsecase 用户数据操作类
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase 构造
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// FindByUserName  通过username查找用户
func (uc *UserUsecase) FindByUserName(ctx context.Context, username string) (*ent.User, error) {
	return uc.repo.GetUserInfo(ctx, username)
}
