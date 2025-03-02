package services

import (
	"Gateway/internal/api/checkout"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
)

type CartItm struct {
	ProductId int64   `json:"product_id"`
	Price     float32 `json:"price"`
	Quantity  int32   `json:"quantity"`
}

type CheckoutReq struct {
	CartItems []*CartItm `json:"cart_item"`
}

func (c *CmsAPP) Checkout(ctx *gin.Context) {
	//tmp, state := ctx.Get("user_id")
	//var userId = tmp.(int64)
	//if !state {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "session is not exist"})
	//	return
	//}
	var req CheckoutReq
	log.Infof("收到结算请求 %+v", req)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 调用微服务
	cartItems := make([]*checkout.CartItem, len(req.CartItems))
	for i, item := range req.CartItems {
		cartItems[i] = &checkout.CartItem{
			ProductId: uint64(item.ProductId),
			Price:     item.Price,
			Quantity:  uint32(item.Quantity),
		}
	}
	log.Infof("结算微服务请求 %+v", cartItems)
	rsp, err := c.checkoutServiceClient.Checkout(ctx, &checkout.CheckoutReq{
		CartItems: cartItems,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(rsp.Products) != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": rsp,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "product not found",
		})
	}

}
