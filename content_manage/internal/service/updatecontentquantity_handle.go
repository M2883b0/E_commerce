package service

import (
	"content_manage/api/operate"
	"context"
)

func (a *AppService) UpdateContentQuantity(ctx context.Context,
	req *operate.UpdateContentQuantityReq) (*operate.UpdateContentQuantityRsp, error) {
	uc := a.uc
	res, err := uc.UpdateContentQuantity(ctx, req.GetId(), req.GetIsAdd(), req.GetQuantity())
	if err != nil {
		return &operate.UpdateContentQuantityRsp{IsSuccess: false}, err
	}
	return &operate.UpdateContentQuantityRsp{IsSuccess: res}, nil
}
