package service

import (
	pb "CartService/api/cart"
	"CartService/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type CartServiceService struct {
	pb.UnimplementedCartServiceServer
	uc *biz.CartUseCase
}

func NewCartServiceService(uc *biz.CartUseCase) *CartServiceService {
	return &CartServiceService{uc: uc}
}

func (s *CartServiceService) AddItem(ctx context.Context, req *pb.AddItemReq) (*pb.AddItemResp, error) {
	item := biz.CartItem{
		UserID:    uint64(req.GetUserId()),
		ProductID: uint64(req.GetItem().GetProductId()),
		Quantity:  uint64(req.GetItem().GetQuantity()),
	}

	if s.uc.IsExist(ctx, &item) {
		cartItems, _, err := s.uc.FindCartItem(ctx, &biz.FindParams{
			UserId:    item.UserID,
			ProductId: item.ProductID,
			Page:      0,
			PageSize:  1,
		})
		item.Quantity += cartItems[0].Quantity
		err = s.uc.UpdateCartItem(ctx, &item)
		if err != nil {
			return &pb.AddItemResp{
				State: false,
			}, nil
		}
	}
	err := s.uc.CreateCartItem(ctx, &item)
	if err != nil {
		return &pb.AddItemResp{State: false}, err
	}

	return &pb.AddItemResp{
		State: true,
	}, nil
}

func (s *CartServiceService) UpdateItem(ctx context.Context, req *pb.UpdateItemReq) (*pb.UpdateItemResp, error) {
	item := biz.CartItem{
		UserID:    uint64(req.GetUserId()),
		ProductID: uint64(req.GetItem().GetProductId()),
		Quantity:  uint64(req.GetItem().GetQuantity()),
	}
	if !s.uc.IsExist(ctx, &item) {
		return &pb.UpdateItemResp{
			State: false,
		}, nil
	}

	if item.Quantity == 0 {
		err := s.uc.DeleteCartItem(ctx, &item)
		if err != nil {
			return &pb.UpdateItemResp{State: false}, err
		}
		return &pb.UpdateItemResp{
			State: true,
		}, nil
	}
	err := s.uc.UpdateCartItem(ctx, &item)
	if err != nil {
		return &pb.UpdateItemResp{State: false}, err
	}

	return &pb.UpdateItemResp{
		State: true,
	}, nil
}
func (s *CartServiceService) GetCart(ctx context.Context, req *pb.GetCartReq) (*pb.GetCartResp, error) {
	findParams := biz.FindParams{
		UserId:    uint64(req.GetUserId()),
		ProductId: 0,
		Page:      0,
		PageSize:  10000000,
	}
	cartItems, total, _ := s.uc.FindCartItem(ctx, &findParams)
	log.Infof("GetCart: %+v", findParams)
	if total == 0 {
		return &pb.GetCartResp{
			Total: 0,
		}, nil
	}
	var ids []int64
	for _, id := range cartItems {
		ids = append(ids, int64(id.ProductID))
	}
	contentInfos, err := s.uc.GetContentInfoByIds(ctx, ids)
	if err != nil {
		return &pb.GetCartResp{
			Total: 0,
		}, nil
	}

	var cartItemsResp []*pb.CartItem
	for i, contentInfo := range contentInfos {
		cartItemsResp = append(cartItemsResp, &pb.CartItem{
			ProductId:      uint32(contentInfo.Id),
			Quantity:       int32(cartItems[i].Quantity),
			Title:          contentInfo.Title,
			Description:    contentInfo.Description,
			PictureUrl:     contentInfo.PictureUrl,
			Price:          contentInfo.Price,
			StoredQuantity: contentInfo.ServerQuantity,
			Categories:     contentInfo.Categories,
		})
	}
	log.Infof("GetCart: %+v", cartItemsResp[0])
	return &pb.GetCartResp{
		Total: uint64(total),
		Items: cartItemsResp,
	}, nil
}
func (s *CartServiceService) EmptyCart(ctx context.Context, req *pb.EmptyCartReq) (*pb.EmptyCartResp, error) {
	state, err := s.uc.DeleteAll(ctx, uint64(req.GetUserId()))
	if err != nil || state == false {
		return &pb.EmptyCartResp{State: false}, nil
	}
	return &pb.EmptyCartResp{
		State: true,
	}, nil
}
