package services

import (
	ai "Gateway/internal/api/ai_agent"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
)

type AIAgentReq struct {
	UserMessage string `json:"user_message"`
}

func (c *CmsAPP) AIAgent(ctx *gin.Context) {
	var req AIAgentReq
	log.Infof("begin AiAgent, user Message is %+v ", req)
	tmp, state := ctx.Get("user_id")
	var userId = tmp.(int64)
	if !state {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "session is not exist"})
		return
	}

	//下面不走，直接db的方法(dao层)，走的是微服务grpc的方法。【内容网关功能很干净了，不走db的操作，转发给grpc去执行操作】
	rsp, err := c.aiAgentClient.UserRequest(ctx, &ai.UserRequestReq{
		UserId:      userId,
		UserMessage: req.UserMessage,
	})
	log.Infof("already begin aiagent server  %+v ", req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": rsp,
	})
}
