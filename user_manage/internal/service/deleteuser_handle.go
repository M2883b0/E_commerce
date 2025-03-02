package service

import (
	"context"
	"user_manage/api/operate"
)

func (a *UserService) DeleteUser(ctx context.Context, req *operate.DeleteUserRequest) (*operate.DeleteUserReply, error) {
	code, msg, err := a.uc.DeleteUser(ctx, req.GetId())
	if err != nil {
		return &operate.DeleteUserReply{
			Code: 400,
			Msg:  msg,
		}, err
	}
	return &operate.DeleteUserReply{
		Code: code,
		Msg:  msg,
	}, nil
}
