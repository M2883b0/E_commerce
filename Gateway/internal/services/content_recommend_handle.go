package services

import (
	"Gateway/internal/api/operate"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 前端的请求数据结构
type ContentRecommendReq struct {
	User_id  int64 `json:"user_id"`   // 用户id
	Page     int32 `json:"page"`      // 页
	PageSize int32 `json:"page_size"` // 页大小
}

func (c *CmsAPP) ContentRecommend(ctx *gin.Context) {
	var req ContentRecommendReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//下面不走，直接db的方法，走的是grpc的方法。【内容网关功能很干净了，不走db的操作，转发给grpc去执行操作】
	rsp, err := c.operateAppClient.RecommendContent(ctx, &operate.RecommendContentReq{
		UserId:   req.User_id,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
			"data": rsp,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": rsp,
	})
}
