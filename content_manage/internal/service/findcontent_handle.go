package service

import (
	"content_manage/api/operate"
	"context"
)

func (a *AppService) FindContent(ctx context.Context,
	req *operate.FindContentReq) (*operate.FindContentRsp, error) {
	//构建请求参数结构
	uc := a.uc
	results, total, err := uc.FindContent(ctx, req.GetQuery(), req.GetPage(), req.GetPageSize()) //调用biz层的实现
	if err != nil {
		return nil, err
	}
	//构建返回的Rsp。FindContentRsp
	var contents []*operate.Content //构建内容
	//用一个for循环，遍历从数据库中拿出来的结果
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
	rsp := &operate.FindContentRsp{ //api的app.proto里面，定义了FindContentRsp结构需要两个内容，total和Contents
		Total:    total,
		Contents: contents,
		Msg:      "执行成功",
		Code:     0,
	}
	return rsp, nil
}
