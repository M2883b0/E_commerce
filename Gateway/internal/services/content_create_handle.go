package services

import (
	"Gateway/internal/api/operate"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ContentCreateReq struct {
	Title       string   `json:"title" binding:"required"` // 内容标题
	Description string   `json:"description"`              // 内容描述
	PictureUrl  string   `json:"picture_url"`
	Price       float32  `json:"price"`
	Quantity    uint32   `json:"quantity"`
	Categories  []string `json:"categories"`
}

type ContentCreateRsp struct {
	Message string `json:"message"`
}

func (c *CmsAPP) ContentCreate(ctx *gin.Context) {
	var req ContentCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//下面不走，直接db的方法(dao层)，走的是微服务grpc的方法。【内容网关功能很干净了，不走db的操作，转发给grpc去执行操作】
	rsp, err := c.operateAppClient.CreateContent(ctx, &operate.CreateContentReq{
		Content: &operate.Content{
			Title:       req.Title,
			Description: req.Description,
			PictureUrl:  req.PictureUrl,
			Price:       req.Price,
			Quantity:    req.Quantity,
			Categories:  req.Categories,
		},
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
