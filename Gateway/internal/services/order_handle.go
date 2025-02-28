package services

import (
	"Gateway/internal/api/order"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PlaceOrderReq struct {
	UserId      uint64       `json:"user_id" binding:"required"`
	PhoneNumber string       `json:"phone_number"`
	Address_    Address      `json:"address"`
	OrderItems  []*OrderItem `json:"order_items"`
}

type Address struct {
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	Country       string `json:"country"`
	ZipCode       uint32 `json:"zip_code"`
}

type OrderItem struct {
	ProductId uint64  `json:"product_id"`
	Quantity  uint32  `json:"quantity"`
	Cost      float32 `json:"cost"`
}

func (c *CmsAPP) PlaceOrder(ctx *gin.Context) {
	var req PlaceOrderReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 调用微服务
	var orderItems []*order.OrderItem
	for _, item := range req.OrderItems {
		orderItems = append(orderItems, &order.OrderItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			Cost:      item.Cost,
		})
	}
	rsp, err := c.orderServiceClient.PlaceOrder(ctx, &order.PlaceOrderReq{
		UserId: req.UserId,
		Address: &order.Address{
			StreetAddress: req.Address_.StreetAddress,
			City:          req.Address_.City,
			Country:       req.Address_.Country,
			ZipCode:       req.Address_.ZipCode,
		},
		PhoneNumber: req.PhoneNumber,
		OrderItems:  orderItems,
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

type ListOrderReq struct {
	UserId uint64 `json:"user_id" binding:"required"` // 内容标题
}

func (c *CmsAPP) ListOrder(ctx *gin.Context) {
	var req ListOrderReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//下面不走，直接db的方法(dao层)，走的是微服务grpc的方法。【内容网关功能很干净了，不走db的操作，转发给grpc去执行操作】
	rsp, err := c.orderServiceClient.ListOrder(ctx, &order.ListOrderReq{
		UserId: req.UserId,
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
