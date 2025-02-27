package service

import (
	"content_manage/api/operate"
	"context"
)

func (a *AppService) RecommendContent(ctx context.Context, req *operate.RecommendContentReq) (*operate.RecommendContentRsp, error) {
	results, total, err := a.uc.RecommendContent(ctx, req.GetUserId(), req.GetPage(), req.GetPageSize())
	if err != nil {
		return nil, err
	}
	var contents []*operate.Content
	for _, r := range results {
		contents = append(contents, &operate.Content{
			Id:          r.ID,
			Title:       r.Title,
			Description: r.Description,
			PictureUrl:  r.Picture_url,
			Price:       r.Price,
			Quantity:    r.Quantity,
			Categories:  r.Categories,
		})
	}
	return &operate.RecommendContentRsp{
		Total:    total,
		Contents: contents,
	}, nil
}
