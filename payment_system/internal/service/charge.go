package service

import (
	"context"
	"payment_system/api/payment"
)

func (p *PaymentService) Charge(ctx context.Context, req *payment.ChargeReq) (*payment.ChargeResp, error) {
	// 支付对外统一charge接口，根据支付方式调用不同支付接口
	//{
	//	switch req.PaymentMethod {
	//	//case payment.PaymentMethod_ALIPAY:
	//		return p.handleAlipayCharge(ctx, req)
	//	//case payment.PaymentMethod_WECHAT_PAY:
	//	//	return p.handleWechatCharge(ctx, req)
	//	default:
	//		return nil, status.Error(codes.InvalidArgument, "unsupported payment method")
	//	}
	//}
	return p.handleAlipayCharge(ctx, req)
}
