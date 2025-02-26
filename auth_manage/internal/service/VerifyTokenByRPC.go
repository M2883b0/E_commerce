package service

import (
	"auth_manage/api/operate"
	"auth_manage/internal/biz"
	"context"
)

func (a *AuthService) VerifyTokenByRPC(ctx context.Context, req *operate.VerifyTokenReq) (*operate.VerifyResp, error) {
	token := req.GetToken()
	uc := a.uc
	res, msg, err := uc.Check_Token(ctx, &biz.Auth{User_id: token})
	if err != nil {
		return nil, err
	}
	return &operate.VerifyResp{
		Res: res,
		Msg: msg,
	}, nil
}

//执行protoc的返回
