package service

import (
	"checkout_system/api/checkout"
	"checkout_system/internal/biz"
	"context"
)

// Checkout implements Checkout.CheckOutServer.
func (c *CheckoutService) Checkout(ctx context.Context, req *checkout.CheckoutReq) (*checkout.CheckoutResp, error) {
	//TODO implement me
	// 默认运费6元
	shippingFee := float32(6)

	// 将 checkout.CartItem 转换为 biz.CartItem
	bizCartItems := make([]*biz.CartItem, len(req.GetCartItems()))
	for i, item := range req.GetCartItems() {
		bizCartItems[i] = &biz.CartItem{
			ProductId:    item.GetProductId(),
			BuyerNeedNum: item.GetQuantity(),
			Price:        item.GetPrice(),
		}
	}
	//深拷贝一份原始信息
	originalBizCartItems := make([]*biz.CartItem, len(req.GetCartItems()))
	for i, item := range req.GetCartItems() {
		bizCartItems[i] = &biz.CartItem{
			ProductId:    item.ProductId,
			BuyerNeedNum: item.Quantity,
			Price:        item.Price,
		}
	}
	// 1.先获取商品信息
	getLatestProductsRsp, err := c.checkoutUc.GetLatestProducts(ctx, &biz.CheckoutPreviewReq{
		bizCartItems,
	})
	// bizCartItems只浅拷贝了
	// 判断价格是否发生了改变

	checkPriceRsp, err := c.checkoutUc.CheckPrice(ctx, &biz.CheckPrice{
		LatestCartItems:   getLatestProductsRsp.CartItems,
		OriginalCartItems: originalBizCartItems,
	})

	if err != nil {
		return nil, err
	}
	// 判断库存
	checkStockRsp, err := c.checkoutUc.CheckStock(ctx, &biz.CheckStock{
		bizCartItems,
	})
	if err != nil {
		return nil, err
	}
	// 计算总价
	calculateTotalPriceRsp, err := c.checkoutUc.CalculateTotalPrice(ctx, &biz.CalculateTotalPrice{
		CartItems: checkStockRsp.CartItems,
	})
	if err != nil {
		return nil, err
	}
	// 计算优惠，返回是否免运费，以及实付价格
	calculateDiscountRsp, err := c.checkoutUc.CalculateDiscount(ctx, &biz.CalculateDiscount{
		TotalPrice: calculateTotalPriceRsp.TotalPrice,
		// 运费默认6元
		ShippingFee: shippingFee,
	})
	// 转为checkout.Product
	products := make([]*checkout.Product, len(getLatestProductsRsp.CartItems))
	for i, item := range checkStockRsp.CartItems {
		products[i] = &checkout.Product{
			ProductId:         item.ProductId,
			Name:              item.Name,
			PictureUrl:        item.PictureUrl,
			Price:             item.Price,
			IsStockSufficient: item.IsStockSufficient,
		}
	}
	return &checkout.CheckoutResp{
		Products:       products,
		ShippingFee:    shippingFee,
		IsFreeShipping: calculateDiscountRsp.IsShippingFree,
		HasChanged:     checkPriceRsp.IsChanged,
		TotalPrice:     calculateTotalPriceRsp.TotalPrice,
		ActualPrice:    calculateDiscountRsp.ActualPrice,
	}, nil
}
