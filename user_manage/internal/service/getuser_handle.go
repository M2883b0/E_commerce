package service

import (
	"context"
	"user_manage/api/operate"
	"user_manage/internal/biz"
)

func (a *UserService) GetUser(ctx context.Context, req *operate.GetUserRequest) (*operate.GetUserReply, error) {
	users, err := a.uc.FindUser(ctx, &biz.FindParams{
		ID: req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	var usersInfo *operate.UserInfo
	usersInfo = &operate.UserInfo{
		Id:          users.ID,
		PhoneNumber: users.Phone_number,
		UserName:    users.User_name,
		UserType:    users.User_type,
		ImgUrl:      users.Img_url,
		Description: users.Description,
	}
	return &operate.GetUserReply{
		User: usersInfo,
	}, nil

}
