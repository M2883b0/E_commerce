package biz

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

type Order struct {
	OrderId       uint64       `json:"order_id"`
	UserID        uint64       `json:"user_id"`
	PhoneNumber   string       `json:"phone_number"`
	IsPaid        string       `json:"is_paid"`
	StreetAddress string       `json:"street_address"`
	City          string       `json:"city"`
	Country       string       `json:"country"`
	ZipCode       uint32       `json:"zip_code"`
	OrderItems    []*OrderItem `json:"order_items"`
}

type OrderItem struct {
	ProductId uint64
	Quantity  uint32
	Cost      float32
}

type OrderRepo interface {
	Create(context.Context, *Order) error
	Update(context.Context, uint64, *Order) error
	IsExist(context.Context, uint64) (bool, error)
	Delete(context.Context, uint64) error
	Find(context.Context, *FindParams) ([]*Order, int64, error)
	UpdateContentInfo(ctx context.Context, params []*UpdateContentItem) (bool, error)
}

type UpdateContentItem struct {
	ProductId int64
	Quantity  int32
	IsAdd     bool
}

// FindParams 查找的参数
type FindParams struct {
	ID          uint64
	PhoneNumber string
	IsPaid      *bool
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
	fmt.Print(ctx, g)
	uc.log.WithContext(ctx).Infof("CreateOrder: %+v", g)
	return uc.repo.Create(ctx, g)
}

func (uc *OrderUseCase) UpdateOrder(ctx context.Context, g *Order) error {
	uc.log.WithContext(ctx).Infof("UpdateOrder: %+v", g)
	return uc.repo.Update(ctx, g.OrderId, g)
}

func (uc *OrderUseCase) DeleteOrder(ctx context.Context, id uint64) error {
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

func (uc *OrderUseCase) FindOrder(ctx context.Context, params *FindParams) ([]*Order, int64, error) {
	orders, total, err := uc.repo.Find(ctx, params)
	if err != nil {
		return nil, 0, err
	}
	return orders, total, nil
}

func (uc *OrderUseCase) UpdateContent(ctx context.Context, param []*UpdateContentItem) bool {
	state, err := uc.repo.UpdateContentInfo(ctx, param)
	if err != nil {
		return false
	}
	return state
}

//执行组合逻辑
