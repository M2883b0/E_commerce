package services

import (
	"content_system/internal/api/operate"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserCreateReq struct {
	Phone_number string `json:"phone_number"`
	Password     string `json:"password"`
	User_name    string `json:"user_name"`
	User_type    int32  `json:"user_type"`
	Img_url      string `json:"img_url"`
	Description  string `json:"description"`
	Address      string `json:"address"`
}

func (c *CmsAPP) UserCreate(ctx *gin.Context) {
	var req UserCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//密码加密
	hashedPassword, err := encryptPassword(req.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	//下面不走，直接db的方法(dao层)，走的是微服务grpc的方法。【内容网关功能很干净了，不走db的操作，转发给grpc去执行操作】
	rsp, err := c.operateUserClient.CreateUser(ctx, &operate.CreateUserRequest{
		User: &operate.UserInfo{
			PhoneNumber: req.Description,
			Password:    hashedPassword,
			UserName:    req.User_name,
			UserType:    req.User_type,
			ImgUrl:      req.Img_url,
			Description: req.Description,
			Address:     req.Address,
		},
	})
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
