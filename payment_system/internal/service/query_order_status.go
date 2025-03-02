package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"payment_system/api/payment"
	"payment_system/internal/biz"
)

func (p *PaymentService) QueryOrderStatus(ctx context.Context, req *payment.QueryOrderStatusReq) (*payment.QueryOrderStatusResp, error) {
	// 主动查询支付结果
	log.Infof("QueryOrderStatus:%+v", req)
	queryPaymentRsp, err := p.alipayUc.QueryPayment(ctx, p.alipayClient, &biz.QueryPayment{
		OutTradeNo: req.OrderId,
	})
	if err != nil {
		return nil, err
	}
	return &payment.QueryOrderStatusResp{
		OrderId: queryPaymentRsp.OutTradeNo,
		Status:  queryPaymentRsp.Status,
	}, nil
}
