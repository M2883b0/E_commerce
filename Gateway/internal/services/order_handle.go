package services

import (
	"Gateway/internal/api/order"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PlaceOrderReq struct {
	PhoneNumber string       `json:"phone_number"`
	Address_    Address      `json:"address"`
	OrderItems  []*OrderItem `json:"order_items"`
}

type Address struct {
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	Country       string `json:"country"`
	ZipCode       int32  `json:"zip_code"`
}

type OrderItem struct {
	ProductId int64   `json:"product_id"`
	Quantity  int32   `json:"quantity"`
	Cost      float32 `json:"cost"`
}

func (c *CmsAPP) PlaceOrder(ctx *gin.Context) {
	var req PlaceOrderReq
	tmp, state := ctx.Get("user_id")
	var userId = tmp.(int64)
	if !state {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "session is not exist"})
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 调用微服务
	var orderItems []*order.OrderItem
	for _, item := range req.OrderItems {
		orderItems = append(orderItems, &order.OrderItem{
			ProductId: item.ProductId,
			Quantity:  uint32(item.Quantity),
			Cost:      item.Cost,
		})
	}
	rsp, err := c.orderServiceClient.PlaceOrder(ctx, &order.PlaceOrderReq{
		UserId: userId,
		Address: &order.Address{
			StreetAddress: req.Address_.StreetAddress,
			City:          req.Address_.City,
			Country:       req.Address_.Country,
			ZipCode:       uint32(req.Address_.ZipCode),
		},
		PhoneNumber: req.PhoneNumber,
		OrderItems:  orderItems,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if rsp == nil || rsp.GetOrderId() == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "error, create order failed",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": rsp,
		})
	}

}

func (c *CmsAPP) ListOrder(ctx *gin.Context) {
	tmp, state := ctx.Get("user_id")
	var userId = tmp.(int64)
	if !state {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "session is not exist"})
		return
	}
	//下面不走，直接db的方法(dao层)，走的是微服务grpc的方法。【内容网关功能很干净了，不走db的操作，转发给grpc去执行操作】
	rsp, err := c.orderServiceClient.ListOrder(ctx, &order.ListOrderReq{
		UserId: userId,
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

type CancelOrderReq struct {
	OrderId int64 `json:"order_id" binding:"required"` // 内容标题
}

func (c *CmsAPP) CancelOrder(ctx *gin.Context) {
	var req CancelOrderReq
	tmp, state := ctx.Get("user_id")
	var userId = tmp.(int64)
	if !state {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "session is not exist"})
		return
	}
	//log.Infof("================================this is user_id %+v==============================================", userId)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//下面不走，直接db的方法(dao层)，走的是微服务grpc的方法。【内容网关功能很干净了，不走db的操作，转发给grpc去执行操作】
	rsp, err := c.orderServiceClient.MarkOrderCancel(ctx, &order.MarkOrderCancelReq{
		UserId:  userId,
		OrderId: req.OrderId,
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

type GetOrderByIdReq struct {
	OrderId int64 `json:"order_id" binding:"required"` // 内容标题
}

func (c *CmsAPP) GetOrderById(ctx *gin.Context) {
	var req GetOrderByIdReq
	tmp, state := ctx.Get("user_id")
	var userId = tmp.(int64)
	if !state {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "session is not exist"})
		return
	}
	//log.Infof("================================this is user_id %+v==============================================", userId)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//下面不走，直接db的方法(dao层)，走的是微服务grpc的方法。【内容网关功能很干净了，不走db的操作，转发给grpc去执行操作】
	rsp, err := c.orderServiceClient.GetOrderById(ctx, &order.GetOrderByIdReq{
		UserId:  userId,
		OrderId: req.OrderId,
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

type DelOrderByIdReq struct {
	OrderId int64 `json:"order_id" binding:"required"` // 内容标题
}

func (c *CmsAPP) DelOrderById(ctx *gin.Context) {
	var req DelOrderByIdReq
	tmp, state := ctx.Get("user_id")
	var userId = tmp.(int64)
	if !state {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "session is not exist"})
		return
	}
	//log.Infof("================================this is user_id %+v==============================================", userId)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//下面不走，直接db的方法(dao层)，走的是微服务grpc的方法。【内容网关功能很干净了，不走db的操作，转发给grpc去执行操作】
	rsp, err := c.orderServiceClient.DelOrderById(ctx, &order.DelOrderByIdReq{
		UserId:  userId,
		OrderId: req.OrderId,
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
