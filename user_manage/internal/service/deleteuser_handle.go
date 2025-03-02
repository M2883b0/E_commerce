package service

import (
	"context"
	"user_manage/api/operate"
)

func (a *UserService) DeleteUser(ctx context.Context, req *operate.DeleteUserRequest) (*operate.DeleteUserReply, error) {
	code, msg, _ := a.uc.DeleteUser(ctx, req.GetId())
	return &operate.DeleteUserReply{
		Code: code,
		Msg:  msg,
	}, nil
}
