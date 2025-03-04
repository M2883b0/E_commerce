package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Auth struct {
	User_id int64
}

type Verfy struct {
	Token string
}

// GreeterRepo is a Greater repo.
type AuthRepo interface {
	SetToken(context.Context, *Auth) (string, error)
	CheckToken(context.Context, *Verfy) (bool, string, int64, error)
	ExpireToken(context.Context, *Verfy) (bool, string, error)
}

type AuthUsecase struct {
	repo AuthRepo
	log  *log.Helper
}

func NewAuthUsecase(repo AuthRepo, logger log.Logger) *AuthUsecase {
	return &AuthUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AuthUsecase) Set_Token(ctx context.Context, a *Auth) (string, error) {
	return uc.repo.SetToken(ctx, a)
}
func (uc *AuthUsecase) Check_Token(ctx context.Context, a *Verfy) (bool, string, int64, error) {
	return uc.repo.CheckToken(ctx, a)
}
func (uc *AuthUsecase) Expire_Token(ctx context.Context, a *Verfy) (bool, string, error) {
	return uc.repo.ExpireToken(ctx, a)
}

//执行组合逻辑（这里没有组合逻辑，直接下一层）
