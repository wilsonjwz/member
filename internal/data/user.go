package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"user/internal/biz"
	"user/internal/data/ent"
	"user/internal/data/ent/user"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) GetUserInfo(ctx context.Context, username string) (*ent.User, error) {
	platUser, err := r.data.db.User.Query().Where(user.UserName(username)).First(ctx)
	if err != nil {
		return nil, err
	}

	return platUser, nil
}
