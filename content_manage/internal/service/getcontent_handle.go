package service

import (
	"content_manage/api/operate"
	"context"
)

func (a *AppService) GetContent(ctx context.Context, req *operate.GetContentReq) (*operate.GetContentRsp, error) {
	contents, err := a.uc.GetContent(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	var contentInfos []*operate.Content
	for _, content := range contents {
		contentInfos = append(contentInfos, &operate.Content{
			Id:          content.ID,
			Title:       content.Title,
			Description: content.Description,
			PictureUrl:  content.Picture_url,
			Price:       content.Price,
			Quantity:    content.Quantity,
			Categories:  content.Categories,
		})
	}
	return &operate.GetContentRsp{
		Contents: contentInfos,
		Msg:      "执行成功",
		Code:     0,
	}, nil
}
