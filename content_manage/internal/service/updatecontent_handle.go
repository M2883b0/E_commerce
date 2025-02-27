package service

import (
	"content_manage/api/operate"
	"content_manage/internal/biz"
	"context"
)

func (a *AppService) UpdateContent(ctx context.Context,
	req *operate.UpdateContentReq) (*operate.UpdateContentRsp, error) {
	content := req.GetContent()
	uc := a.uc
	err := uc.UpdateContent(ctx, &biz.Content{
		ID:          content.GetId(),
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
	return &operate.UpdateContentRsp{}, nil
}
