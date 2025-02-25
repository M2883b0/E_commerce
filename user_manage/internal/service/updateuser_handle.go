package service

import (
	"context"
	"user_manage/api/operate"
	"user_manage/internal/biz"
)

func (a *UserService) UpdateUser(ctx context.Context, req *operate.UpdateUserRequest) (*operate.UpdateUserReply, error) {
	user := req.GetUser()
	err := a.uc.UpdateUser(ctx, &biz.User{
		ID:           user.Id,
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
	return &operate.UpdateUserReply{}, nil
}
