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
	CartItem CartItem `json:"cart_item"`
}

func (c *CmsAPP) AddCartItem(ctx *gin.Context) {
	tmp, state := ctx.Get("user_id")
	var userId = tmp.(int64)
	if !state {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "session is not exist"})
		return
	}
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
		UserId: userId,
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

type UpdateItemReq struct {
	CartItem CartItem `json:"cart_item"`
}

func (c *CmsAPP) UpdateCartItem(ctx *gin.Context) {
	tmp, state := ctx.Get("user_id")
	var userId = tmp.(int64)
	if !state {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "session is not exist"})
		return
	}
	var req UpdateItemReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 调用微服务
	var cartItem = cart.CartItem{
		ProductId: uint32(req.CartItem.ProductId),
		Quantity:  req.CartItem.Quantity,
	}
	rsp, err := c.cartServiceClient.UpdateItem(ctx, &cart.UpdateItemReq{
		UserId: userId,
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

func (c *CmsAPP) GetCart(ctx *gin.Context) {
	tmp, state := ctx.Get("user_id")
	var userId = tmp.(int64)
	if !state {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "session is not exist"})
		return
	}
	//下面不走，直接db的方法(dao层)，走的是微服务grpc的方法。【内容网关功能很干净了，不走db的操作，转发给grpc去执行操作】
	rsp, err := c.cartServiceClient.GetCart(ctx, &cart.GetCartReq{
		UserId: userId,
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

func (c *CmsAPP) EmptyCart(ctx *gin.Context) {
	tmp, state := ctx.Get("user_id")
	var userId = tmp.(int64)
	if !state {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "session is not exist"})
		return
	}
	//下面不走，直接db的方法(dao层)，走的是微服务grpc的方法。【内容网关功能很干净了，不走db的操作，转发给grpc去执行操作】
	rsp, err := c.cartServiceClient.EmptyCart(ctx, &cart.EmptyCartReq{
		UserId: userId,
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
