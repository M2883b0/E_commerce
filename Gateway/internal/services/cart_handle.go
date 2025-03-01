package services

import (
	"Gateway/internal/api/cart"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CartItem struct {
	ProductId int64 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

type AddItemReq struct {
	UserId   int64    `json:"user_id"`
	CartItem CartItem `json:"cart_item"`
}

func (c *CmsAPP) AddItem(ctx *gin.Context) {
	var req AddItemReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 调用微服务
	var cartItem = cart.CartItem{
		ProductId: uint32(req.CartItem.ProductId),
		Quantity:  req.CartItem.Quantity,
	}
	rsp, err := c.cartServiceClient.AddItem(ctx, &cart.AddItemReq{
		UserId: uint32(req.UserId),
		Item:   &cartItem,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if rsp.GetState() {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": rsp,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "add cart item failure!",
		})
	}

}

// GetCartReq get cart 接口。获取购物车信息
type GetCartReq struct {
	UserId int64 `json:"user_id" binding:"required"` // 内容标题
}

func (c *CmsAPP) GetCart(ctx *gin.Context) {
	var req GetCartReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//下面不走，直接db的方法(dao层)，走的是微服务grpc的方法。【内容网关功能很干净了，不走db的操作，转发给grpc去执行操作】
	rsp, err := c.cartServiceClient.GetCart(ctx, &cart.GetCartReq{
		UserId: uint32(req.UserId),
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if rsp.GetTotal() != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": rsp,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "cart is empty,or something wrong",
		})
	}
}

type EmptyCartReq struct {
	UserId int64 `json:"user_id" binding:"required"` // 内容标题
}

func (c *CmsAPP) EmptyCart(ctx *gin.Context) {
	var req EmptyCartReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//下面不走，直接db的方法(dao层)，走的是微服务grpc的方法。【内容网关功能很干净了，不走db的操作，转发给grpc去执行操作】
	rsp, err := c.cartServiceClient.EmptyCart(ctx, &cart.EmptyCartReq{
		UserId: uint32(req.UserId),
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if rsp.GetState() {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": rsp,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "empty cart failed",
		})
	}
}
