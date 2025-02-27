package service

import (
	"auth_manage/api/operate"
	"auth_manage/internal/biz"
)

type AuthService struct {
	operate.UnimplementedAuthServer
	uc *biz.AuthUsecase
}

func NewAuthService(uc *biz.AuthUsecase) *AuthService {

	return &AuthService{uc: uc}
}
