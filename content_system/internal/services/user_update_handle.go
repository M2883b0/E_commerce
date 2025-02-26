package services

import (
	"content_system/internal/api/operate"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 前端更新，需要传的内容结构。id是必须传的，需要知道对哪条内容进行更新操作
type UserUpdateReq struct {
	ID           int64  `json:"id"`
	Phone_number string `json:"phone_number"`
	Password     string `json:"password"`
	User_name    string `json:"user_name"`
	User_type    int32  `json:"user_type"`
	Img_url      string `json:"img_url"`
	Description  string `json:"description"`
	Address      string `json:"address"`
}

func (c *CmsAPP) UserUpdate(ctx *gin.Context) {
	var req UserUpdateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.ID == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "更新操作，需要指定更新的ID"})
	}

	//密码加密
	hashedPassword, err := encryptPassword(req.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if req.Password == "" { //如果不更新密码
		hashedPassword = "" //加密后，强行置为空
	}
	//下面不走，直接db的方法，走的是grpc的方法。【内容网关功能很干净了，不走db的操作，转发给grpc去执行操作】
	rsp, err := c.operateUserClient.UpdateUser(ctx, &operate.UpdateUserRequest{
		User: &operate.UserInfo{
			Id:          req.ID,
			PhoneNumber: req.Phone_number,
			Password:    hashedPassword,
			UserName:    req.User_name,
			UserType:    req.User_type,
			ImgUrl:      req.Img_url,
			Description: req.Description,
			Address:     req.Address,
		},
	})
	if err != nil { //出现错误
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "更新成功",
		"data": rsp,
	})
}
