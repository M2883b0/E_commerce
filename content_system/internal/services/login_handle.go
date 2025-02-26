package services

import (
	"content_system/internal/api/operate"
	"content_system/internal/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	if rsp.Code == 0 { //如果没报错，就生成sessionID
		sessionID, err = cms.GenerateSessionId(context.Background(), rsp.User.PhoneNumber)
	}

	//上面是Session的方法，在GenerateSessionId函数中，需要使用redis存入内存中。
	//这里我们使用jwt的方法，加密的方法 import "content_system/jwt"
	//sessionID, err := jwt.SetToken(req.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "系统设置session错误，请重新尝试"})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":        rsp.Msg,
		"code":       rsp.Code,
		"data":       rsp.User,
		"session_id": sessionID,
	})

	////初始化dao层的实例，用dao层的方法，实现功能逻辑
	//accountDao := dao.NewAccountDao(cms.db)
	//
	////先判断数据库中是否存在这个用户
	//account, err := accountDao.GetInfoByUserID(req.Phone_number)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "账号不存在，请先注册"})
	//	return
	//}
	////如果存在用户，则比较数据库中的密码和用户传的密码，是否一致
	//if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password)); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "密码不正确"})
	//	return
	//}
	//
	////账号密码校验成功，接下来返回Session信息给前端
	//sessionID, err := cms.GenerateSessionId(context.Background(), account.Phone_number)
	//
	////上面是Session的方法，在GenerateSessionId函数中，需要使用redis存入内存中。
	////这里我们使用jwt的方法，加密的方法 import "content_system/jwt"
	////sessionID, err := jwt.SetToken(req.UserID)
	//
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "系统设置session错误，请重新尝试"})
	//}
	//
	////回包
	//c.JSON(http.StatusOK, gin.H{ //登录成功
	//	"msg": "ok",
	//	"data": &LoginRes{
	//		SessionID:   sessionID,
	//		ID:          int64(account.ID),
	//		user_name:   account.User_name,
	//		user_type:   account.User_type,
	//		img_url:     account.Img_url,
	//		description: account.Description,
	//	},
	//})
	//
	//return
}

func (cms *CmsAPP) GenerateSessionId(ctx context.Context, userID string) (string, error) {
	//SessionID的生成
	session_id := uuid.New().String()
	//SessionID的持久化。用于判断用户的session_id是否正确
	key := utils.GetSessionKey(userID)
	err := cms.rdb.Set(ctx, key, session_id, session_time).Err()
	if err != nil { //验证用户给的session_id是否有效
		fmt.Printf("redis set sessionid err:%v\n", err)
		return "", err
	}
	//同时以session_id为key，再存储一个。用于判断session_id是否过期
	authkey := utils.GetAuthKey(session_id)
	err = cms.rdb.Set(ctx, authkey, time.Now().Unix(), session_time).Err()
	if err != nil { //判断session_id是否过期了
		fmt.Printf("redis set sessionid err:%v\n", err)
		return "", err
	}
	return session_id, nil
}
