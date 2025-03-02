package data

import (
	"ai_manage/internal/biz"
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
