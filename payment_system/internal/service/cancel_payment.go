package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"payment_system/api/payment"
	"payment_system/internal/biz"
)

func (p *PaymentService) Cancel(ctx context.Context, req *payment.CancelReq) (*payment.CancelResp, error) {
	log.Infof("取消支付:%d", req.OrderId)
	// 1.初始化alipay  在service构造函数中进行了初始化
	// 2.调用alipay取消支付
	cancelPaymentRsp, err := p.alipayUc.CancelPayment(ctx, p.alipayClient, &biz.CancelReq{
		OutTradeNo: req.OrderId,
	})

	if err != nil {
		return nil, err
	}
	return &payment.CancelResp{
		OrderId: cancelPaymentRsp.OutTradeNo,
		Msg:     cancelPaymentRsp.Msg,
		Code:    cancelPaymentRsp.Code,
	}, nil

}
