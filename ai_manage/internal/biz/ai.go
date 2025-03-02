package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Req struct {
	User_id int64
	Query   string
}
type Reply struct {
	Reply string
	Code  int32
	Msg   string
}

type AiRepo interface {
	ChatAI(context.Context, *Req) (*Reply, error)
}

type AiUsecase struct {
	repo AiRepo
	log  *log.Helper
}

func NewAiUsecase(repo AiRepo, logger log.Logger) *AiUsecase {
	return &AiUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AiUsecase) Chat_AI(ctx context.Context, u *Req) (*Reply, error) {
	return uc.repo.ChatAI(ctx, u)
}
