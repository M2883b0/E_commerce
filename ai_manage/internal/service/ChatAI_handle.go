package service

import (
	"ai_manage/api/operate"
	"context"
)

func (a *AiService) ChatAI_Handle(ctx context.Context, req *operate.ChatAIRequest) (reply *operate.ChatAIReply, err error) {
	return &operate.ChatAIReply{
		Code:  0,
		Msg:   "success",
		Reply: "你好",
	}, nil
}
