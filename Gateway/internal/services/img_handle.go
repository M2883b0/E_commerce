package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// 模型的绑定与验证：绑定：【处理我们自定义的数据结构体】；字段验证：【不能为空，手机号规则，邮箱规则,长度】
// 请求结构体
type HelloReq struct {
	UserId string `json:"user_id" binding:"required"` //定义前端传过来，必须要包含name字段(required)
}

// 响应结构体
type HelloRes struct {
	Message string `json:"message" binding:"required"` //定义后端的响应数据
}

func (cms *CmsAPP) Hello(c *gin.Context) {
	var req HelloReq
	//要求前端按照HelloReq结构体中定义的那样，如果不按照，则会报错
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 获取上传的文件
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	// 定义保存文件的路径和文件名
	uploadPath := "./uploads" // 你可以根据需要更改这个路径
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fileName := fmt.Sprintf("%s.jpg", req.UserId)
	filePath := filepath.Join(uploadPath, fileName)

	// 创建文件
	out, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer out.Close()

	// 将上传的文件内容复制到新创建的文件中
	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//如果前端正确发送请求，则执行下面的程序，返回响应数据
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": &HelloRes{
			Message: "头像更新完成",
		},
	})
}
