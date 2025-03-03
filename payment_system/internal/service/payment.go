package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/smartwalle/alipay/v3"
	"os"
	"payment_system/api/payment"
	"payment_system/internal/biz"
)

// PaymentService is a Payment service.
type PaymentService struct {
	payment.UnimplementedPaymentServiceServer
	alipayUc     *biz.AlipayUsecase
	alipayClient *alipay.Client
}

// NewPaymentService new a greeter service.
func NewPaymentService(alipayUc *biz.AlipayUsecase) *PaymentService {
	appId := os.Getenv("APP_ID")
	if appId != "" {
		log.Infof("app_id: %+v", appId)
	} else {
		log.Infof("app_id为空")
		return nil
	}
	privateKey := os.Getenv("PRIVATE_KEY")
	if privateKey != "" {
		log.Infof("private_key: %+v", privateKey)
	} else {
		log.Infof("private_key为空")
		return nil
	}
	alipayPublicKey := os.Getenv("ALIPAY_PUBLIC_KEY")
	if alipayPublicKey != "" {
		log.Infof("alipay_public_key: %+v", alipayPublicKey)
	} else {
		log.Infof("alipay_public_key为空")
		return nil
	}

	// 传入的数值，转为字符串
	isProductionString := os.Getenv("IS_PRODUCTION")
	if isProductionString != "" {
		log.Infof("is_production: %+v", isProductionString)
	} else {
		log.Infof("is_production为空")
		return nil
	}
	var isProduction bool
	if isProductionString == "0" {
		isProduction = false
	} else if isProductionString == "1" {
		isProduction = true
	}
	var client, err = alipay.New(appId, privateKey, isProduction)
	if err != nil {
		log.Errorf("初始化支付宝失败: %+v", err)
		panic(err)
	}

	// 加载支付宝公钥
	err = client.LoadAliPayPublicKey(alipayPublicKey)
	if err != nil {
		panic(err)
	}
	return &PaymentService{
		alipayUc:     alipayUc,
		alipayClient: client}
}
