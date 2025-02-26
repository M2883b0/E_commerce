package service

import (
	"auth_manage/api/operate"
	"auth_manage/internal/biz"
	"context"
)

func (a *AuthService) DeliverTokenByRPC(ctx context.Context, req *operate.DeliverTokenReq) (*operate.DeliveryResp, error) {
	user_id := req.GetUserId()
	uc := a.uc
	token, err := uc.Set_Token(ctx, &biz.Auth{User_id: user_id})
	if err != nil {
		return nil, err
	}
	return &operate.DeliveryResp{Token: token}, nil
}

//执行protoc的返回
