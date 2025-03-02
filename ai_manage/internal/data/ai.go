package data

import (
	"ai_manage/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type aiRepo struct {
	data *Data
	log  *log.Helper
}

func NewAiRepo(data *Data, logger log.Logger) biz.AiRepo {
	return &aiRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *aiRepo) ChatAI(ctx context.Context, a *biz.Req) (*biz.Reply, error) {

	return &biz.Reply{
		Code:  0,
		Msg:   "success",
		Reply: "你好",
	}, nil
}
