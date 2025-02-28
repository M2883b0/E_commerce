package service

import (
	"content_manage/api/operate"
	"content_manage/internal/biz"
	"context"
)

func (a *AppService) UpdateContentQuantity(ctx context.Context,
	req *operate.UpdateContentQuantityReq) (*operate.UpdateContentQuantityRsp, error) {
	uc := a.uc
	content := req.GetQuantityReq()
	var quantity_list []*biz.QuantityDetail
	for _, v := range content {
		quantity_list = append(quantity_list, &biz.QuantityDetail{
			ID:       v.Id,
			Is_add:   v.IsAdd,
			Quantity: v.Quantity,
		})
	}
	res, err := uc.UpdateContentQuantity(ctx, quantity_list)
	if err != nil {
		return &operate.UpdateContentQuantityRsp{IsSuccess: false}, err
	}
	return &operate.UpdateContentQuantityRsp{IsSuccess: res}, nil
}
