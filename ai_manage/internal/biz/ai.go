package biz

import "github.com/go-kratos/kratos/v2/log"

type AiRepo interface {
}

type AiUsecase struct {
	repo AiRepo
	log  *log.Helper
}

func NewAiUsecase(repo AiRepo, logger log.Logger) *AiUsecase {
	return &AiUsecase{repo: repo, log: log.NewHelper(logger)}
}
