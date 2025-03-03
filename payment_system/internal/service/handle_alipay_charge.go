package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"payment_system/api/payment"
	"payment_system/internal/biz"
)

func (p *PaymentService) handleAlipayCharge(ctx context.Context, req *payment.ChargeReq) (*payment.ChargeResp, error) {
	log.Infof("支付宝支付:%d", req.OrderId)
	// 1.初始化alipay  在service构造函数中进行了初始化
	// 2.调用alipay支付
	alipayTradePreCreateRsp, err := p.alipayUc.Trade(ctx, p.alipayClient, &biz.TradeReq{
		OutTradeNo: req.GetOrderId(),
		Subject:    req.Subject,
	})

	if err != nil {
		return nil, err
	}

	return &payment.ChargeResp{
		OrderId: alipayTradePreCreateRsp.OutTradeNo,
		QrUrl:   alipayTradePreCreateRsp.QrCode,
	}, err

}
