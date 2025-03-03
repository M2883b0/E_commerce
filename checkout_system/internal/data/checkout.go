package data

import (
	"checkout_system/api/operate"
	"checkout_system/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type checkoutRepo struct {
	productClient *ProductClient
	log           *log.Helper
}

// NewGreeterRepo .
func NewCheckoutRepo(productClient *ProductClient, logger log.Logger) biz.CheckoutRepo {
	return &checkoutRepo{
		productClient: productClient,
		log:           log.NewHelper(logger),
	}
}

func (c *checkoutRepo) FindCartItem(ctx context.Context, checkout *biz.CheckoutPreviewReq) (*biz.GetLatestProductsRsp, error) {
	// 更新购物车里的东西的信息
	log.Infof("查找购物车里的东西的信息:%+v", checkout)
	idList := make([]int64, len(checkout.CartItems))
	for i, cartItem := range checkout.CartItems {
		idList[i] = int64(cartItem.ProductId)
	}
	productService := c.productClient.client
	products, err := productService.GetContent(ctx,
		&operate.GetContentReq{
			Id: idList,
		})

	if products == nil {
		log.Errorf("购物车商品不存在")
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	mapLatestProduct := make(map[int64]*operate.Content, len(products.Contents))
	for _, product := range products.Contents {
		mapLatestProduct[product.Id] = product
	}
	for _, cartItem := range checkout.CartItems {
		latestCartItem := mapLatestProduct[int64(cartItem.ProductId)]
		cartItem = &biz.CartItem{
			ProductId:  uint64(latestCartItem.GetId()),
			Name:       latestCartItem.GetTitle(),
			PictureUrl: latestCartItem.GetPictureUrl(),
			Price:      float32(latestCartItem.GetPrice()),
			Stock:      latestCartItem.GetQuantity(),
		}
	}
	return (*biz.GetLatestProductsRsp)(checkout), nil
}
