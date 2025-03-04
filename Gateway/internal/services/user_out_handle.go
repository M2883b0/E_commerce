package services

import (
	"Gateway/internal/api/operate"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SessionKey = "session_id" //鉴权的id，前端需要在Headers中加入这个参数session_id
)

func (c *CmsAPP) UserOut(ctx *gin.Context) {
	sessionID := ctx.GetHeader(SessionKey)
	//下面不走，直接db的方法，走的是grpc的方法。【内容网关功能很干净了，不走db的操作，转发给grpc去执行操作】
	rsp, err := c.operateAuthClient.ExpireTokenByRPC(ctx, &operate.ExpireTokenReq{Token: sessionID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 400,
			"msg":  err.Error(),
			"data": rsp,
		})
		return
	}
	//code := 0
	//msg := "ok"
	//if rsp == nil || rsp.User.Id == 0 {
	//	code = 400
	//	msg = "用户不存在"
	//}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  rsp.Msg,
		"data": rsp,
	})
}
