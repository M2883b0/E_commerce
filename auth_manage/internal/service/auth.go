package service

import (
	"auth_manage/api/operate"
	"auth_manage/internal/biz"
)

type AuthService struct {
	operate.UnimplementedAuthServer
	uc *biz.GreeterUsecase
}

func NewAuthService(uc *biz.GreeterUsecase) *AuthService {
	return &AuthService{uc: uc}
}
