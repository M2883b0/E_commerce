package service

import (
	"context"
	"user_manage/api/operate"
	"user_manage/internal/biz"
)

func (a *UserService) CreateUser(ctx context.Context, req *operate.CreateUserRequest) (*operate.CreateUserReply, error) {
	user := req.GetUser()
	err := a.uc.CreateUser(ctx, &biz.User{
		Phone_number: user.GetPhoneNumber(),
		Password:     user.GetPassword(),
		User_name:    user.GetUserName(),
		User_type:    user.GetUserType(),
		Img_url:      user.GetImgUrl(),
		Description:  user.GetDescription(),
	})
	if err != nil {
		return nil, err
	}
	return &operate.CreateUserReply{}, nil
}
