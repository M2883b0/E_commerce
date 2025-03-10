package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"payment_system/api/order"
	"payment_system/internal/biz"
)

type OrderStatusRepo struct {
	orderClient *OrderClient
	log         *log.Helper
}

// NewOrderStatusRepo .
func NewOrderStatusRepo(client *OrderClient, logger log.Logger) biz.OrderStatusRepo {
	return &OrderStatusRepo{
		orderClient: client,
		log:         log.NewHelper(logger),
	}
}

func (c *OrderStatusRepo) MarkOrderPaid(ctx context.Context, orderId int64) (*order.MarkOrderPaidResp, error) {
	log.Infof("标记订单为已支付:%+v", orderId)
	markOrderPaidRsp, err := c.orderClient.client.MarkOrderPaid(ctx, &order.MarkOrderPaidReq{OrderId: int64(orderId)})
	if err != nil {
		return nil, err
	}
	return markOrderPaidRsp, nil
}

func (c *OrderStatusRepo) MarkOrderCancel(ctx context.Context, orderId int64) (*order.MarkOrderCancelResp, error) {
	log.Infof("标记订单取消支付:%+v", orderId)
	markOrderCancelRsp, err := c.orderClient.client.MarkOrderCancel(ctx, &order.MarkOrderCancelReq{OrderId: int64(orderId)})
	if err != nil {
		return nil, err
	}
	return markOrderCancelRsp, nil
}

func (c *OrderStatusRepo) GetOrderInfo(ctx context.Context, orderId int64) (*order.GetOrderByIdResp, error) {
	log.Infof("获取订单信息:%+v", orderId)
	orderInfo, err := c.orderClient.client.GetOrderById(ctx, &order.GetOrderByIdReq{
		UserId:  0,
		OrderId: orderId,
	})
	if err != nil {
		return nil, err
	}
	return orderInfo, nil
}
