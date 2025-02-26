package services

import (
	"content_system/internal/api/operate"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const session_time = 8 * time.Hour

// 模型的绑定与验证：绑定：【处理我们自定义的数据结构体】；字段验证：【不能为空，手机号规则，邮箱规则,长度】
// 前端请求的结构体数据【登录只需要前端传给我账号和密码】
type LoginReq struct {
	Phone_number string `json:"phone_number" binding:"required"` //定义前端传过来，必须要包含该字段(required)
	Password     string `json:"password" binding:"required"`
}

// 响应结构体
type LoginRes struct {
	SessionID   string `json:"session_id"` //定义返回内容，需要返回一个sessionID
	ID          int64  `json:"id"`         //同时返回一些用户的信息
	user_name   string `json:"user_name"`
	user_type   int32  `json:"user_type"`
	img_url     string `json:"img_url"`
	description string `json:"description"`
	address     string `json:"address"`
}

func (cms *CmsAPP) Login(c *gin.Context) {
	//初始化前端的返回结构实例
	var req LoginReq

	//要求前端按照HelloReq结构体中定义的那样，如果不按照，则会报错
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rsp, err := cms.operateUserClient.Login(c, &operate.LoginRequest{
		Login: &operate.LoginInfo{
			PhoneNumber: req.Phone_number,
			Password:    req.Password,
		},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sessionID := ""
	if rsp.Code == 0 { //如果没报错,登录成功，就生成sessionID
		jwt_rsp, err := cms.operateAuthClient.DeliverTokenByRPC(c, &operate.DeliverTokenReq{UserId: rsp.User.PhoneNumber})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		sessionID = jwt_rsp.Token
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "系统设置session错误，请重新尝试"})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":        rsp.Msg,
		"code":       rsp.Code,
		"data":       rsp.User,
		"session_id": sessionID,
	})
}
