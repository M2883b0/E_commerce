package service

import (
	"context"
	"user_manage/api/operate"
	"user_manage/internal/biz"
)

func (a *UserService) Register(ctx context.Context, req *operate.RegisterRequest) (*operate.RegisterReply, error) {
	user := req.GetRegister()
	res, err := a.uc.RegisterUser(ctx, &biz.Register{
		Phone_number: user.GetPhoneNumber(),
		Password:     user.GetPassword(),
		User_name:    user.GetUserName(),
	})
	var reg *operate.RegisterReply
	reg = &operate.RegisterReply{
		Code: res.Code,
		Msg:  res.Msg,
	}
	return reg, err
}
