package biz

import (
	"context"
	"fmt"
	"net/url"
	"payment_system/api/order"
	//"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/smartwalle/alipay/v3"
)

//var (
//	// ErrUserNotFound is user not found.
//	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
//)

// Payment is a Payment model.
type Payment struct {
	OrderID int64
	Amount  string
	Status  string
	QrUrl   string
}

type TradeReq struct {
	OutTradeNo  int64   `json:"out_trade_no"`
	TotalAmount float32 `json:"total_amount"`
	Subject     string
	ProductCode string
}
type TradeRsp struct {
	OutTradeNo int64 `json:"out_trade_no"`
	QrCode     string
}

type QueryPayment struct {
	OutTradeNo int64 `json:"out_trade_no"`
}

type QueryPaymentRsp struct {
	OutTradeNo int64 `json:"out_trade_no"`
	Status     string
}

type Notify struct {
	Params map[string][]string
}

type NotifyRsp struct {
	State string
}

const (
	TradeStatusWaitBuyerPay TradeStatus = "WAIT_BUYER_PAY" //（交易创建，等待买家付款）
	TradeStatusClosed       TradeStatus = "TRADE_CLOSED"   //（未付款交易超时关闭，或支付完成后全额退款）
	TradeStatusSuccess      TradeStatus = "TRADE_SUCCESS"  //（交易支付成功）
	TradeStatusFinished     TradeStatus = "TRADE_FINISHED"
)

type TradeStatus string

// PaymentRepo is a Payment repo.
type PaymentRepo interface {
	// 创建支付订单
	Create(ctx context.Context, payment *Payment) error
	// 更新支付订单
	Update(ctx context.Context, payment *Payment) error
	// 查询支付订单
	FindByID(ctx context.Context, orderId int64) (*Payment, error)
	//ListByHello(context.Context, string) ([]*Greeter, error)
	//ListAll(context.Context) ([]*Greeter, error)
}

type OrderStatusRepo interface {
	// 更新订单状态为已支付
	MarkOrderPaid(ctx context.Context, orderId int64) (*order.MarkOrderPaidResp, error)
	// 更新订单为取消支付
	MarkOrderCancel(ctx context.Context, orderId int64) (*order.MarkOrderCancelResp, error)
	// 获取订单信息
	GetOrderInfo(ctx context.Context, orderId int64) (*order.GetOrderByIdResp, error)
}

// GreeterUsecase is a Greeter usecase.
type AlipayUsecase struct {
	paymentRepo     PaymentRepo
	orderStatusRepo OrderStatusRepo
	log             *log.Helper
}

// NewAlipayUsecase 创建支付宝实例
func NewAlipayUsecase(payRepo PaymentRepo, orderStatusRepo OrderStatusRepo, logger log.Logger) *AlipayUsecase {
	return &AlipayUsecase{paymentRepo: payRepo, orderStatusRepo: orderStatusRepo, log: log.NewHelper(logger)}
}

func (uc *AlipayUsecase) Trade(ctx context.Context, client *alipay.Client, req *TradeReq) (*TradeRsp, error) {
	uc.log.WithContext(ctx).Infof("调用支付宝接口: %+v", req)
	orderInfo, err := uc.orderStatusRepo.GetOrderInfo(ctx, req.OutTradeNo)
	if err != nil {
		return nil, err
	}
	if orderInfo.GetOrder().GetUserId() == 0 {
		return &TradeRsp{
			OutTradeNo: 0,
			QrCode:     "",
		}, nil
	}
	totalAmount := orderInfo.GetOrder().ActualPayment
	totalAmountString := fmt.Sprintf("%.2f", totalAmount)

	// 先查询数据库，如果订单是等待支付状态，返回QrUrl
	payment, err := uc.paymentRepo.FindByID(ctx, req.OutTradeNo)
	if err != nil {
		return nil, err
	}
	if payment != nil && payment.Status == "WAIT_BUYER_PAY" {
		return &TradeRsp{
			OutTradeNo: req.OutTradeNo,
			QrCode:     payment.QrUrl,
		}, nil
	}
	if payment != nil && payment.Status == "TRADE_SUCCESS" {
		return &TradeRsp{
			OutTradeNo: req.OutTradeNo,
			QrCode:     "订单已支付",
		}, nil
	}
	if payment != nil && payment.Status == "TRADE_FINISHED" {
		return &TradeRsp{
			OutTradeNo: req.OutTradeNo,
			QrCode:     "订单已支付",
		}, nil
	}
	if payment != nil && payment.Status == "TRADE_CLOSED" {
		return &TradeRsp{
			OutTradeNo: req.OutTradeNo,
			QrCode:     "订单已关闭",
		}, nil
	}
	//
	trade := alipay.Trade{
		Subject:     req.Subject,
		OutTradeNo:  fmt.Sprintf("%d", req.OutTradeNo),
		TotalAmount: totalAmountString,
		ProductCode: "FAST_INSTANT_TRADE_PAY",
		// 30分钟支付超时
		TimeoutExpress: "30m",
		// 异步通知url
		//NotifyURL:     "http://127.0.0.1:8080/payment/notify",
	}
	log.Infof("支付参数:%+v", trade)

	tradePagePay := alipay.TradePagePay{
		Trade:     trade,
		QRPayMode: "0",
		//【描述】PC扫码支付的方式。
		//支持前置模式和跳转模式。
		//前置模式是将二维码前置到商户的订单确认页的模式。需要商户在自己的页面中以 iframe 方式请求支付宝页面。具体支持的枚举值有以下几种：
		//0：订单码-简约前置模式，对应 iframe 宽度不能小于600px，高度不能小于300px；
		//1：订单码-前置模式，对应iframe 宽度不能小于 300px，高度不能小于600px；
		//3：订单码-迷你前置模式，对应 iframe 宽度不能小于 75px，高度不能小于75px；
		//4：订单码-可定义宽度的嵌入式二维码，商户可根据需要设定二维码的大小。
		//跳转模式下，用户的扫码界面是由支付宝生成的，不在商户的域名下。
	}

	//log.Infof("支付:%+v",tradePagePay)

	var url2 *url.URL
	url2, err = client.TradePagePay(tradePagePay)
	if err != nil {
		fmt.Println(err)
	}

	// 创建支付数据
	if err := uc.paymentRepo.Create(ctx, &Payment{
		OrderID: req.OutTradeNo,
		Amount:  totalAmountString,
		Status:  "WAIT_BUYER_PAY",
		QrUrl:   url2.String(),
	}); err != nil {
		return nil, err
	}

	log.Infof("url2: %+v", url2)

	return &TradeRsp{
		OutTradeNo: req.OutTradeNo,
		QrCode:     url2.String(),
	}, nil
	//alipay.TradePreCreateRsp{
	//	Error:      alipay.Error{},
	//	OutTradeNo: "", // 商户订单号
	//	QRCode:     "", // 支付二维码链接
	//}
}

// QueryPayment 查询订单状态
func (uc *AlipayUsecase) QueryPayment(ctx context.Context, client *alipay.Client, req *QueryPayment) (*QueryPaymentRsp, error) {
	uc.log.WithContext(ctx).Infof("查询支付状态: %+v", req)
	tradeQuery := alipay.TradeQuery{
		OutTradeNo: fmt.Sprintf("%d", req.OutTradeNo),
	}
	log.Infof("查询支付状态:%+v", tradeQuery)
	// 查询数据库，如果订单是已成功支付，支付关闭，支付完成不可退款
	payment, err := uc.paymentRepo.FindByID(ctx, req.OutTradeNo)
	if err != nil {
		return nil, err
	}
	if payment == nil {
		return nil, fmt.Errorf("订单不存在")
	}
	tradeStatus := TradeStatus(payment.Status)
	if tradeStatus == TradeStatusSuccess || tradeStatus == TradeStatusClosed || tradeStatus == TradeStatusFinished {
		return &QueryPaymentRsp{
			OutTradeNo: req.OutTradeNo,
			Status:     string(tradeStatus),
		}, nil
	}
	// 调用支付宝查询订单状态
	tradeRsp, err := client.TradeQuery(ctx, tradeQuery)
	if err != nil {
		return nil, err
	}
	log.Infof("支付宝返回订单状态：%+v", tradeRsp)
	if tradeRsp.TradeStatus == alipay.TradeStatusSuccess {
		// 支付成功
		// Todo调用订单服务更新订单状态

		markOrderPaidRsp, err := uc.orderStatusRepo.MarkOrderPaid(ctx, req.OutTradeNo)
		if err != nil {
			return nil, err
		}
		// Todo 更新支付订单状态
		err = uc.paymentRepo.Update(ctx, &Payment{
			OrderID: req.OutTradeNo,
			Status:  string(tradeRsp.TradeStatus),
		})
		if err != nil {
			return nil, err
		}

		if !markOrderPaidRsp.State {
			return nil, fmt.Errorf("订单状态更新失败")
		}
	}
	if tradeRsp.TradeStatus == alipay.TradeStatusClosed {
		// 支付超时
		// Todo调用订单服务更新订单状态
		markOrderCancelRsp, err := uc.orderStatusRepo.MarkOrderCancel(ctx, req.OutTradeNo)
		if err != nil {
			return nil, err
		}
		if !markOrderCancelRsp.State {
			return nil, fmt.Errorf("订单状态更新失败")
		}
		// Todo 更新支付订单状态
		err = uc.paymentRepo.Update(ctx, &Payment{
			OrderID: req.OutTradeNo,
			Status:  string(tradeRsp.TradeStatus),
		})
		if err != nil {
			return nil, err
		}
	}

	return &QueryPaymentRsp{
		OutTradeNo: req.OutTradeNo,
		Status:     string(tradeRsp.TradeStatus),
	}, nil
}

type CancelReq struct {
	OutTradeNo int64 `json:"out_trade_no"`
}
type CancelResp struct {
	OutTradeNo int64 `json:"out_trade_no"`
	Code       string
	Msg        string
}

// 取消支付
func (uc *AlipayUsecase) CancelPayment(ctx context.Context, client *alipay.Client, req *CancelReq) (*CancelResp, error) {
	uc.log.WithContext(ctx).Infof("CancelPayment: %+v", req)
	tradeCancel := alipay.TradeCancel{
		OutTradeNo: fmt.Sprintf("%d", req.OutTradeNo),
	}
	tradeRsp, err := client.TradeCancel(ctx, tradeCancel)
	// 通知订单服务关闭订单
	markOrderCancelState, err := uc.orderStatusRepo.MarkOrderCancel(ctx, req.OutTradeNo)
	if err != nil {
		return nil, err
	}
	if !markOrderCancelState.State {
		return nil, fmt.Errorf("订单状态更新失败")
	}
	// 通知支付订单状态
	err = uc.paymentRepo.Update(ctx, &Payment{
		OrderID: req.OutTradeNo,
		Status:  string(alipay.TradeStatusClosed),
	})
	if err != nil {
		return nil, err
	}

	//const (
	//	CodeSuccess                  Code = "10000" // 接口调用成功
	//	CodeOrderSuccessPayInProcess Code = "10003" // 表示订单创建成功，支付处理中
	//	CodeUnknowError              Code = "20000" // 服务不可用
	//	CodeInvalidAuthToken         Code = "20001" // 授权权限不足
	//	CodeMissingParam             Code = "40001" // 缺少必选参数
	//	CodeInvalidParam             Code = "40002" // 非法的参数
	//	CodeInsufficientConditions   Code = "40003" // 条件异常
	//	CodeBusinessFailed           Code = "40004" // 业务处理失败
	//	CodeCallLimited              Code = "40005" // 调用频次超限
	//	CodePermissionDenied         Code = "40006" // 权限不足
	//)

	//type Error struct {
	//	Code    Code   `json:"code"`
	//	Msg     string `json:"msg"`
	//	SubCode string `json:"sub_code"`
	//	SubMsg  string `json:"sub_msg"`
	//}

	var msg string
	//判断支付宝公共状态码
	if tradeRsp.Error.Code == alipay.CodeBusinessFailed {
		msg = "BusinessFailed"
	}
	if tradeRsp.Error.Code == alipay.CodeSuccess {
		msg = "Success"
	}

	return &CancelResp{
		OutTradeNo: req.OutTradeNo,
		Msg:        msg,
		Code:       string(tradeRsp.Error.Code),
	}, err
}

//func (uc *AlipayUsecase) NotifyHandle(ctx context.Context, client *alipay.Client, req *Notify) (*NotifyResp, error) {
//	uc.log.WithContext(ctx).Infof("异步返回处理: %+v", req)
//	// 验签
//	err := client.VerifySign(req.Params)
//	// 验签失败
//	if err != nil {
//		return &NotifyRsp{State: "fail"}, err
//	}
//	// 验签成功
//	notification, err := client.DecodeNotification(req.Params)
//	if err != nil {
//		return nil, err
//	}
//	outTradeNo, err := strconv.Atoi(notification.OutTradeNo)
//	if err != nil {
//		return nil, err
//	}
//	switch notification.TradeStatus {
//	case alipay.TradeStatusSuccess:
//		log.Infof("支付成功%+v", notification.OutTradeNo)
//		payment, err := uc.paymentRepo.FindByID(ctx, int64(outTradeNo))
//		if payment.Status == string(alipay.TradeStatusSuccess) {
//			return &NotifyRsp{
//				State: "success",
//			}, nil
//		}
//		markOrderPaidRsp, err := uc.orderStatusRepo.MarkOrderPaid(ctx, int64(outTradeNo))
//		if err != nil {
//			return nil, err
//		}
//		if !markOrderPaidRsp.State {
//			return nil, fmt.Errorf("订单状态更新失败")
//		}
//		err = uc.paymentRepo.Update(ctx, &Payment{
//			OrderID: int64(outTradeNo),
//			Status:  string(alipay.TradeStatusSuccess),
//		})
//	case alipay.TradeStatusClosed:
//		log.Infof("支付关闭%+v", notification.OutTradeNo)
//		markOrderCancelRsp, err := uc.orderStatusRepo.MarkOrderCancel(ctx, int64(outTradeNo))
//		if err != nil {
//			return nil, err
//		}
//		if !markOrderCancelRsp.State {
//			return nil, fmt.Errorf("订单状态更新失败")
//		}
//		err = uc.paymentRepo.Update(ctx, &Payment{
//			OrderID: int64(outTradeNo),
//			Status:  string(alipay.TradeStatusClosed),
//		})
//	}
//	return nil, nil
//}
