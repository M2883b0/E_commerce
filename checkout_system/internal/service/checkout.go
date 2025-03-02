package service

import (
	"checkout_system/api/checkout"
	"checkout_system/internal/biz"
)

type CheckoutService struct {
	checkout.UnimplementedCheckoutServiceServer
	checkoutUc *biz.CheckoutUsecase
}

// NewCheckoutService new a Checkout service.
func NewCheckoutService(checkoutUc *biz.CheckoutUsecase) *CheckoutService {
	return &CheckoutService{
		checkoutUc: checkoutUc,
	}
}
