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

	for _, cartItem := range checkout.CartItems {
		// 添加打印语句
		c.log.Infof("传进来的 cart item: ProductID=%d, Name=%s, Price=%.2f",
			cartItem.ProductId,
			cartItem.Name,
			cartItem.Price)
	}

	idList := make([]int64, len(checkout.CartItems))
	for i, cartItem := range checkout.CartItems {
		idList[i] = int64(cartItem.ProductId)
	}

	productService := c.productClient.client
	products, err := productService.GetContent(ctx,
		&operate.GetContentReq{
			Id: idList,
		})

	for _, item := range products.Contents {
		log.Infof("商品接口获得的信息: ID=%d, Title=%s, Price=%.2f, Quantity=%d",
			item.Id,
			item.Title,
			item.Price,
			item.Quantity)
	}

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
	for i, cartItem := range checkout.CartItems {
		latestCartItem := mapLatestProduct[int64(cartItem.ProductId)]

		log.Infof("Map中的商品信息: ID=%d, Title=%s, Price=%.2f, Quantity=%d",
			latestCartItem.Id,
			latestCartItem.Title,
			latestCartItem.Price,
			latestCartItem.Quantity)

		cartItem = &biz.CartItem{
			ProductId:  uint64(latestCartItem.GetId()),
			Name:       latestCartItem.GetTitle(),
			PictureUrl: latestCartItem.GetPictureUrl(),
			Price:      float32(latestCartItem.GetPrice()),
			Stock:      latestCartItem.GetQuantity(),
		}
		checkout.CartItems[i] = cartItem

	}

	for _, cartItem := range checkout.CartItems {
		// 添加打印语句
		// 添加完整字段的打印语句
		c.log.Infof("更新后的 cart item: ProductID=%d, Name=%s, Price=%.2f, Stock=%d, Picture=%s",
			cartItem.ProductId,
			cartItem.Name,
			cartItem.Price,
			cartItem.Stock,      // 新增库存字段
			cartItem.PictureUrl) // 新增图片字段
	}

	return (*biz.GetLatestProductsRsp)(checkout), nil
}
