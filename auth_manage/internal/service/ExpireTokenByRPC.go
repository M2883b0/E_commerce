package service

import (
	"auth_manage/api/operate"
	"auth_manage/internal/biz"
	"context"
)

func (a *AuthService) ExpireTokenByRPC(ctx context.Context, req *operate.ExpireTokenReq) (*operate.ExpireTokenResp, error) {
	uc := a.uc
	res, msg, err := uc.Expire_Token(ctx, &biz.Verfy{
		Token: req.GetToken(),
	})
	if err != nil {
		return nil, err
	}
	return &operate.ExpireTokenResp{
		Res: res,
		Msg: msg,
	}, nil
}
