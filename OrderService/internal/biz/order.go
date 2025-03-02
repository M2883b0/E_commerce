package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

type Order struct {
	OrderId        int64        `json:"order_id"`
	UserID         int64        `json:"user_id"`
	OriginalCharge float32      `json:"original_charge"`
	ActualPayment  float32      `json:"actual_payment"`
	IsFreeShipping bool         `json:"is_free_shipping"`
	ShippingFee    float32      `json:"shipping_fee"`
	PhoneNumber    string       `json:"phone_number"`
	OrderState     string       `json:"order_state"`
	StreetAddress  string       `json:"street_address"`
	City           string       `json:"city"`
	Country        string       `json:"country"`
	ZipCode        uint32       `json:"zip_code"`
	OrderItems     []*OrderItem `json:"order_items"`
}

type OrderItem struct {
	ProductId int64
	Quantity  uint32
	Cost      float32
}

type OrderRepo interface {
	Create(context.Context, *Order) error
	Update(context.Context, int64, *Order) error
	IsExist(context.Context, int64) (bool, error)
	Delete(context.Context, int64) error
	Find(context.Context, *FindParams) ([]*Order, int64, error)
	UpdateContentInfo(ctx context.Context, params []*UpdateContentItem) (bool, error)
	CheckoutOrder(ctx context.Context, params []*CheckoutOrderItem) (*CheckoutResp, error)
}

// 更新商品微服务的参数
type UpdateContentItem struct {
	ProductId int64
	Quantity  int32
	IsAdd     bool
}

// 调用结算微服务的参数
type CheckoutOrderItem struct {
	ProductId int64
	Price     float32
	Quantity  int32
}

type CheckoutResp struct {
	ActualPrice    float32
	TotalPrice     float32
	IsFreeShipping bool
	ShippingFee    float32
	HasChanged     bool
}

// FindParams 查找的参数
type FindParams struct {
	UserID      int64
	OrderId     int64
	PhoneNumber string
	OrderState  string
	Page        uint32
	PageSize    uint32
}

type OrderUseCase struct {
	repo OrderRepo
	log  *log.Helper
}

func NewOrderUseCase(repo OrderRepo, logger log.Logger) *OrderUseCase {
	return &OrderUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *OrderUseCase) CreateOrder(ctx context.Context, g *Order) error {
	uc.log.WithContext(ctx).Infof("CreateOrder: %+v", g)
	return uc.repo.Create(ctx, g)
}

func (uc *OrderUseCase) UpdateOrder(ctx context.Context, g *Order) error {
	uc.log.WithContext(ctx).Infof("UpdateOrder: %+v", g)
	return uc.repo.Update(ctx, g.OrderId, g)
}

func (uc *OrderUseCase) DeleteOrder(ctx context.Context, id int64) error {
	uc.log.WithContext(ctx).Infof("DeleteOrder: %+v", id)
	//复合操作，现判断该用户是否存在，再执行删除操作
	ok, err := uc.repo.IsExist(ctx, id)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("user no exist")
	}
	//用户存在的情况,执行删除操作
	return uc.repo.Delete(ctx, id)
}

func (uc *OrderUseCase) FindOrderByUserId(ctx context.Context, UserId int64) ([]*Order, int64) {
	uc.log.WithContext(ctx).Infof("FindOrder: %+v", UserId)
	orders, total, err := uc.repo.Find(ctx, &FindParams{
		UserID:      UserId,
		OrderId:     0,
		PhoneNumber: "",
		OrderState:  "",
		Page:        0,
		PageSize:    9999999,
	})
	if err != nil {
		return nil, 0
	}
	return orders, total
}

func (uc *OrderUseCase) FindOrderById(ctx context.Context, UserId int64, OrderId int64) Order {
	uc.log.WithContext(ctx).Infof("FindOrderById: userId %+v, OrderId %+v", UserId, OrderId)
	orders, _, err := uc.repo.Find(ctx, &FindParams{
		UserID:      UserId,
		OrderId:     OrderId,
		PhoneNumber: "",
		OrderState:  "",
		Page:        0,
		PageSize:    0,
	})
	if err != nil {
		return Order{UserID: -1}
	}
	return *orders[0]
}

func (uc *OrderUseCase) UpdateContent(ctx context.Context, param []*UpdateContentItem) bool {
	state, err := uc.repo.UpdateContentInfo(ctx, param)
	if err != nil {
		return false
	}
	return state
}

func (uc *OrderUseCase) CheckoutOrder(ctx context.Context, param []*CheckoutOrderItem) (*CheckoutResp, error) {
	return uc.repo.CheckoutOrder(ctx, param)
}

//执行组合逻辑
