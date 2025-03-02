package service

import (
	"ai_manage/api/operate"
	"ai_manage/internal/biz"
)

type AiService struct {
	operate.UnimplementedAiServer

	uc *biz.AiUsecase
}

func NewAiService(uc *biz.AiUsecase) *AiService {
	return &AiService{uc: uc}
}
