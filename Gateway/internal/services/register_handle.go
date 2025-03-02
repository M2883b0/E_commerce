package services

import (
	"Gateway/internal/api/operate"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 模型的绑定与验证：绑定：【处理我们自定义的数据结构体】；字段验证：【不能为空，手机号规则，邮箱规则,长度】
// 前端请求的结构体数据集【注册账号，前端要传给我：用户名、密码、名称这3个信息】
type RegisterReq struct {
	Phone_number string `json:"phone_number" binding:"required"` //定义前端传过来，必须要包含该字段(required)
	Password     string `json:"password" binding:"required"`
	User_name    string `json:"user_name" binding:"required"`
}

// 响应结构体
type RegisterRes struct {
	Message string `json:"message" binding:"required"` //定义返回内容，返回一个json消息
}

func (cms *CmsAPP) Register(c *gin.Context) {
	//初始化前端的返回结构实例
	var req RegisterReq

	//要求前端按照RegisterReq结构体中定义的那样，如果不按照，则会报错
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"前端传参error": err.Error()})
		return
	}

	//如果前端正确发送请求，则执行下面的程序，返回响应数据
	//fmt.Printf("register info = %+v \n", req) //%v是只打印【值】，%+v是打印【字段+值】

	////实现密码加密
	//hashedPassword, err := encryptPassword(req.Password)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//}
	rsp, err := cms.operateUserClient.Register(c, &operate.RegisterRequest{
		Register: &operate.RegisterInfo{
			PhoneNumber: req.Phone_number,
			Password:    req.Password,
			UserName:    req.User_name,
		},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  rsp.Msg,
		"code": rsp.Code,
	})
}

// 密码加密函数
//func encryptPassword(password string) (string, error) {
//	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//	if error != nil {
//		fmt.Println("Error hashing password:", error)
//		return "", error
//	}
//	return string(hashedPassword), nil
//}
