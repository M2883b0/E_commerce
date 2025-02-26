package services

import (
	"content_system/internal/api/operate"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 前端的请求数据结构
type UserFindReq struct {
	ID int64 `json:"id"`
}

func (c *CmsAPP) UserFind(ctx *gin.Context) {
	var req UserFindReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//下面不走，直接db的方法，走的是grpc的方法。【内容网关功能很干净了，不走db的操作，转发给grpc去执行操作】
	rsp, err := c.operateUserClient.GetUser(ctx, &operate.GetUserRequest{Id: req.ID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	code := 0
	msg := "ok"
	if rsp == nil || rsp.User.Id == 0 {
		code = 400
		msg = "用户不存在"
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": rsp,
	})
}
