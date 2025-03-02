package service

import (
	v1 "ai_manage/api/helloworld/v1"
	"ai_manage/internal/biz"
)

type AiService struct {
	v1.UnimplementedAiServer

	uc *biz.AiUsecase
}

func NewAiService(uc *biz.AiUsecase) *AiService {
	return &AiService{uc: uc}
}
