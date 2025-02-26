package services

import (
	"content_system/internal/api/operate"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 前端需要传的参数：id是必须的，因为需要知道删除的是哪个数据
type UserDeleteReq struct {
	ID int64 `json:"id" binding:"required"` // ID
}

func (c *CmsAPP) UserDelete(ctx *gin.Context) {
	var req UserDeleteReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//下面不走，直接db的方法，走的是grpc的方法。【内容网关功能很干净了，不走db的操作，转发给grpc去执行操作】
	rsp, err := c.operateUserClient.DeleteUser(ctx, &operate.DeleteUserRequest{Id: req.ID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 400,
			"msg":  err.Error(),
			"data": rsp,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "删除成功",
		"data": rsp,
	})
}
