package services

import (
	"Gateway/internal/api/payment"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CancelReq struct {
	OrderId int64 `json:"product_id"`
}
type ChargeReq struct {
	OrderId int64 `json:"product_id"`
}

type QueryOrderStatusReq struct {
	OrderId int64 `json:"product_id"`
}

func (c *CmsAPP) Charge(ctx *gin.Context) {
	var req payment.ChargeReq
	// 调用微服务
	rsp, err := c.paymentServiceClient.Charge(ctx, &payment.ChargeReq{
		OrderId:        req.OrderId,
		Subject:        req.Subject,
		PaymentMethod:  "",
		IdempotencyKey: "",
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
	var req payment.QueryOrderStatusReq
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
	// 调用微服务
	rsp, err := c.paymentServiceClient.Cancel(ctx, &payment.CancelReq{
		OrderId: req.OrderId,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": rsp.Code,
		"msg":  rsp.Msg,
		"data": rsp,
	})
}
