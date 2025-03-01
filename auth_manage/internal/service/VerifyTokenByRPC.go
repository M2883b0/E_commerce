package service

import (
	"auth_manage/api/operate"
	"auth_manage/internal/biz"
	"context"
)

func (a *AuthService) VerifyTokenByRPC(ctx context.Context, req *operate.VerifyTokenReq) (*operate.VerifyResp, error) {
	uc := a.uc
	res, msg, userId, err := uc.Check_Token(ctx, &biz.Verfy{
		Token: req.GetToken(),
	})
	if err != nil {
		return nil, err
	}
	return &operate.VerifyResp{
		Res:    res,
		Msg:    msg,
		UserId: userId,
	}, nil
}

//执行protoc的返回
