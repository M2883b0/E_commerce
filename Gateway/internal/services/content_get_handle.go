package services

import (
	"Gateway/internal/api/operate"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 前端的请求数据结构
type ContentGetReq struct {
	ID int64 `json:"id"` // 内容ID
}

func (c *CmsAPP) ContentGet(ctx *gin.Context) {
	var req ContentGetReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//下面不走，直接db的方法，走的是grpc的方法。【内容网关功能很干净了，不走db的操作，转发给grpc去执行操作】
	rsp, err := c.operateAppClient.GetContent(ctx, &operate.GetContentReq{
		Id: req.ID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": rsp,
	})
}
