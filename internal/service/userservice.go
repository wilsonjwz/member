package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"user/internal/biz"
	"user/internal/pkg/util"

	pb "user/api/user/v1"
)

type UserServiceService struct {
	pb.UnimplementedUserServiceServer
	log  *log.Helper
	user *biz.UserUsecase
}

func NewUserServiceService(user *biz.UserUsecase, logger log.Logger) *UserServiceService {
	return &UserServiceService{
		user: user,
		log:  log.NewHelper(logger),
	}
}

func (s *UserServiceService) ClientManualAuthorize(ctx context.Context, req *pb.ManualAuthorizeRequest) (*pb.ManualAuthorizeReply, error) {
	s.log.Infof("input data %v", req)
	platUser, err := s.user.FindByUserName(ctx, req.Username)
	if err != nil {
		s.log.Error(err)
		return nil, errors.NotFound(pb.ErrorReason_USER_NOT_FOUND.String(), "error")
	}
	s.log.Infof("plat_user data %v", platUser)
	// 检查密码
	isRight, err := util.CheckPassword(req.Password, platUser.Password, platUser.PasswordStr)
	if err != nil || isRight == false {
		s.log.Error(err)
		return nil, errors.BadRequest(pb.ErrorReason_WRONG_PASSWORD.String(), "帐号或密码错误")
	}

	return &pb.ManualAuthorizeReply{
		User: &pb.User{
			PlatUserId: int64(platUser.ID),
			Username:   platUser.UserName,
			Telnum:     platUser.TelNum,
		},
	}, nil
}
