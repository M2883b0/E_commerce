package service

import (
	"context"
	"user_manage/api/operate"
	"user_manage/internal/biz"
)

func (a *UserService) Login(ctx context.Context, req *operate.LoginRequest) (*operate.LoginReply, error) {
	user := req.GetRegister()
	res, err := a.uc.LoginUser(ctx, &biz.Login{
		Phone_number: user.GetPhoneNumber(),
		Password:     user.GetUserName(),
	})
	if err != nil {
		return nil, err
	}
	var usersInfo *operate.UserInfo
	usersInfo = &operate.UserInfo{
		Id:          res.User.ID,
		PhoneNumber: res.User.Phone_number,
		UserName:    res.User.User_name,
		UserType:    res.User.User_type,
		ImgUrl:      res.User.Img_url,
		Description: res.User.Description,
	}
	return &operate.LoginReply{
		Code: res.Code,
		Msg:  res.Msg,
		User: usersInfo,
	}, nil
}
