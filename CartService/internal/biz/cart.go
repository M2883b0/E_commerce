package biz

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

type CartItem struct {
	UserID    uint64
	ProductID uint64
	Quantity  uint64
}

type ContentInfo struct {
	Id             int64
	Title          string
	Description    string
	PictureUrl     string
	Price          float32
	ServerQuantity uint32
	Categories     []string
}

type CartRepo interface {
	Create(context.Context, *CartItem) error
	Update(context.Context, *CartItem) error
	IsExist(context.Context, *CartItem) bool
	Delete(context.Context, *CartItem) error
	FindCartByUserId(context.Context, *FindParams) ([]*CartItem, int64, error)
	GetContentInfoById(context.Context, uint64) (*ContentInfo, error)
}

type UpdateContentItem struct {
	ProductId int64
	Quantity  int32
	IsAdd     bool
}

type CartUseCase struct {
	repo CartRepo
	log  *log.Helper
}

func NewCartUseCase(repo CartRepo, logger log.Logger) *CartUseCase {
	return &CartUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CartUseCase) IsExist(ctx context.Context, g *CartItem) bool {
	fmt.Print(ctx, g)
	return uc.repo.IsExist(ctx, g)
}

func (uc *CartUseCase) CreateCartItem(ctx context.Context, g *CartItem) error {
	uc.log.WithContext(ctx).Infof("CreateCartItem: %+v", g)
	if uc.repo.IsExist(ctx, g) {
		return uc.UpdateCartItem(ctx, g)
	}
	return uc.repo.Create(ctx, g)
}

func (uc *CartUseCase) UpdateCartItem(ctx context.Context, g *CartItem) error {
	uc.log.WithContext(ctx).Infof("UpdateCartItem: %+v", g)
	return uc.repo.Update(ctx, g)
}

func (uc *CartUseCase) DeleteCartItem(ctx context.Context, g *CartItem) error {
	ok := uc.repo.IsExist(ctx, g)
	if !ok {
		return errors.New("cart item no exist")
	}
	//用户存在的情况,执行删除操作
	return uc.repo.Delete(ctx, g)
}

type FindParams struct {
	UserId    uint64
	ProductId uint64
	Page      uint32
	PageSize  uint64
}

func (uc *CartUseCase) DeleteAll(ctx context.Context, userId uint64) (bool, error) {
	uc.log.WithContext(ctx).Infof("DeleteAll: %+v", userId)
	cartItems, _, err := uc.FindCartItem(ctx, &FindParams{
		UserId:    userId,
		ProductId: 0,
		Page:      0,
		PageSize:  99999999999999,
	})
	if err != nil {
		return false, err
	}
	for _, cartItem := range cartItems {
		if err = uc.repo.Delete(ctx, cartItem); err != nil {
			return false, err
		}
	}
	return true, nil
}

func (uc *CartUseCase) FindCartItem(ctx context.Context, param *FindParams) ([]*CartItem, int64, error) {
	uc.log.WithContext(ctx).Infof("FindCartItem: %+v", param)
	cartItems, total, err := uc.repo.FindCartByUserId(ctx, param)
	if err != nil {
		return nil, 0, err
	}
	return cartItems, total, nil
}

func (uc *CartUseCase) GetContentInfoById(ctx context.Context, id uint64) (*ContentInfo, error) {
	return uc.repo.GetContentInfoById(ctx, id)
}

//执行组合逻辑
