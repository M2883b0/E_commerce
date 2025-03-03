package services

import (
	"Gateway/internal/api/payment"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CancelReq struct {
	OrderId int64 `json:"order_id"`
}
type ChargeReq struct {
	OrderId int64  `json:"order_id"`
	Subject string `json:"subject"`
}

type QueryOrderStatusReq struct {
	OrderId int64 `json:"order_id"`
}

func (c *CmsAPP) Charge(ctx *gin.Context) {
	var req ChargeReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 调用微服务
	rsp, err := c.paymentServiceClient.Charge(ctx,
		&payment.ChargeReq{
			OrderId: req.OrderId,
			Subject: req.Subject,
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

func (c *CmsAPP) QueryOrderStatus(ctx *gin.Context) {
	var req QueryOrderStatusReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rsp, err := c.paymentServiceClient.QueryOrderStatus(ctx, &payment.QueryOrderStatusReq{
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
func (c *CmsAPP) Cancel(ctx *gin.Context) {
	//tmp, state := ctx.Get("user_id")
	//var userId = tmp.(int64)
	//if !state {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "session is not exist"})
	//	return
	//}
	var req CancelReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 调用微服务
	rsp, err := c.paymentServiceClient.Cancel(ctx, &payment.CancelReq{
		OrderId: req.OrderId,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": "0",
		"msg":  "ok",
		"data": rsp,
	})
}
