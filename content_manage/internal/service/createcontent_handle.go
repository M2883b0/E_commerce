package service

import (
	"content_manage/api/operate"
	"content_manage/internal/biz"
	"context"
)

func (a *AppService) CreateContent(ctx context.Context,
	req *operate.CreateContentReq) (*operate.CreateContentRsp, error) {
	content := req.GetContent()
	uc := a.uc
	err := uc.CreateContent(ctx, &biz.Content{
		Title:       content.GetTitle(),
		Description: content.GetDescription(),
		Picture_url: content.GetPictureUrl(),
		Price:       content.GetPrice(),
		Quantity:    content.GetQuantity(),
		Categories:  content.GetCategories(),
	})
	if err != nil {
		return nil, err
	}
	return &operate.CreateContentRsp{
		Code: 0,
		Msg:  "执行成功",
	}, nil
}
