package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

//var (
//	// ErrUserNotFound is user not found.
//	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
//)

type CartItem struct {
	ProductId         uint64
	BuyerNeedNum      uint32
	Price             float32
	Name              string
	PictureUrl        string
	Stock             uint32
	IsStockSufficient bool
}

// Checkout is a Checkout model.
type CheckoutPreviewReq struct {
	CartItems []*CartItem
}

type CheckStock struct {
	CartItems []*CartItem
}

type CheckStockRsp struct {
	CartItems []*CartItem
}

//type ProductInfo struct {
//	ProductId         uint64
//	ProductName       string
//	PictureUrl        string
//	Price             float32
//	Stock             uint64
//	buyerNeed         uint32
//	IsStockSufficient bool
//}

type GetLatestProductsRsp struct {
	CartItems []*CartItem
}

type CheckPrice struct {
	OriginalCartItems []*CartItem
	LatestCartItems   []*CartItem
}
type CheckPriceRsp struct {
	IsChanged bool
}
type CalculateTotalPrice struct {
	CartItems []*CartItem
}

type CalculateTotalPriceRsp struct {
	TotalPrice float32
}

type CalculateDiscount struct {
	ShippingFee float32
	TotalPrice  float32
}

type CalculateDiscountRsp struct {
	ActualPrice    float32
	IsShippingFree bool
}

// CheckoutRepo is a Checkout repo.
type CheckoutRepo interface {
	FindCartItem(ctx context.Context, c *CheckoutPreviewReq) (*GetLatestProductsRsp, error)
}

// CheckoutUsecase is a Greeter usecase.
type CheckoutUsecase struct {
	repo CheckoutRepo
	log  *log.Helper
}

// NewCheckoutUsecase new a Greeter usecase.
func NewCheckoutUsecase(repo CheckoutRepo, logger log.Logger) *CheckoutUsecase {
	return &CheckoutUsecase{repo: repo, log: log.NewHelper(logger)}
}

// 调用商品服务查询商品信息,包括价格，库存等信息
func (uc *CheckoutUsecase) GetLatestProducts(ctx context.Context, req *CheckoutPreviewReq) (*GetLatestProductsRsp, error) {
	log.Infof("开始获取最新商品信息: %+v", req)
	// 查询商品信息Todo 转为productInfo,只留IsStockSufficient不填充
	// 构造商品id列表
	products, err := uc.repo.FindCartItem(ctx, req)
	if products == nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	// 返回productInfo数组
	return products, nil
}

func (uc *CheckoutUsecase) CheckPrice(ctx context.Context, req *CheckPrice) (bool, error) {
	uc.log.WithContext(ctx).Infof("CheckPrice%+v", req)
	// 将 LatestCartItems 转换为 map，键为 ProductId，值为 *CartItem
	latestProductMap := make(map[uint64]*CartItem)
	for _, item := range req.LatestCartItems {
		latestProductMap[item.ProductId] = item
	}
	// 返回商品是否存在改变
	for _, originalCartItem := range req.OriginalCartItems {
		latestCartItem := latestProductMap[originalCartItem.ProductId]
		log.Infof("CheckPrice: %+v", latestCartItem)
		if originalCartItem.Price != latestCartItem.Price {
			return true, nil
		}
	}
	log.Infof("CheckPrice is not change")
	return false, nil
}

func (uc *CheckoutUsecase) CheckStock(ctx context.Context, req *CheckStock) (*CheckStockRsp, error) {
	log.Infof("检查库存: %+v", req)
	// 检查库存,更新IsStockSufficient
	for _, cartItem := range req.CartItems {
		// 查询商品信息
		if cartItem.Stock < cartItem.BuyerNeedNum {
			cartItem.IsStockSufficient = false
		} else {
			cartItem.IsStockSufficient = true
		}
	}
	return &CheckStockRsp{CartItems: req.CartItems}, nil
}

func (uc *CheckoutUsecase) CalculateTotalPrice(ctx context.Context, c *CalculateTotalPrice) (float32, error) {
	log.Infof("计算总价: %+v", c)

	var totalPrice float32
	totalPrice = 0

	// 计算总金额
	for _, cartItem := range c.CartItems {
		fmt.Println("111计算总金额", cartItem)
		productPrice := cartItem.Price
		if cartItem.IsStockSufficient == true {
			// 库存充足才计算总价钱
			totalPrice += productPrice * float32(cartItem.BuyerNeedNum)
		}
	}

	return totalPrice, nil
}

func (uc *CheckoutUsecase) CalculateDiscount(ctx context.Context, req *CalculateDiscount) (*CalculateDiscountRsp, error) {
	log.Infof("计算优惠后的实付: %+v", req)
	totalPrice := req.TotalPrice
	// 金额大于99免基础运费
	var isShippingFree bool
	if req.TotalPrice > 99 {
		isShippingFree = true
	} else {
		isShippingFree = false
		// 需要运费就加运费
		totalPrice += req.ShippingFee
	}
	// 满300-50
	discount := 50
	if req.TotalPrice >= 300 {
		return &CalculateDiscountRsp{
			ActualPrice:    totalPrice - float32(discount),
			IsShippingFree: isShippingFree,
		}, nil
	}
	return &CalculateDiscountRsp{
		ActualPrice:    totalPrice,
		IsShippingFree: isShippingFree,
	}, nil
}
