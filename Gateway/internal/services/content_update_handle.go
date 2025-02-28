package services

import (
	"Gateway/internal/api/operate"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 前端更新，需要传的内容结构。id是必须传的，需要知道对哪条内容进行更新操作
type ContentUpdateReq struct {
	ID          int64    `json:"id" binding:"required"` // 内容标题
	Title       string   `json:"title"`                 // 内容标题
	Description string   `json:"description"`           // 内容描述
	PictureUrl  string   `json:"picture_url"`
	Price       float32  `json:"price"`
	Quantity    uint32   `json:"quantity"`
	Categories  []string `json:"categories"`
}

// 后端返回的响应结构
type ContentUpdateRsp struct {
	Message string `json:"message"`
}

func (c *CmsAPP) ContentUpdate(ctx *gin.Context) {
	var req ContentUpdateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//下面不走，直接db的方法，走的是grpc的方法。【内容网关功能很干净了，不走db的操作，转发给grpc去执行操作】
	rsp, err := c.operateAppClient.UpdateContent(ctx, &operate.UpdateContentReq{
		Content: &operate.Content{
			Id:          req.ID,
			Title:       req.Title,
			Description: req.Description,
			PictureUrl:  req.PictureUrl,
			Price:       req.Price,
			Quantity:    req.Quantity,
			Categories:  req.Categories,
		},
	})
	if err != nil { //出现错误
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": rsp,
	})
}
