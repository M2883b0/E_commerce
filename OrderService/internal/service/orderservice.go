package service

import (
	"OrderService/internal/biz"
	"context"
	"fmt"

	pb "OrderService/api/order"
)

type OrderServiceService struct {
	pb.UnimplementedOrderServiceServer
	uc *biz.OrderUseCase
}

func NewOrderServiceService(uc *biz.OrderUseCase) *OrderServiceService {
	return &OrderServiceService{uc: uc}
}

func (s *OrderServiceService) PlaceOrder(ctx context.Context, req *pb.PlaceOrderReq) (*pb.PlaceOrderResp, error) {

	var orderItemBiz []*biz.OrderItem
	var updateItems []*biz.UpdateContentItem

	for _, orderItemReq := range req.GetOrderItems() {
		orderItemBiz = append(orderItemBiz,
			&biz.OrderItem{
				Quantity:  orderItemReq.GetQuantity(),
				ProductId: orderItemReq.GetProductId(),
				Cost:      orderItemReq.GetCost(),
			})
		updateItems = append(updateItems, &biz.UpdateContentItem{
			ProductId: int64(orderItemReq.GetProductId()),
			Quantity:  int32(orderItemReq.GetQuantity()),
			IsAdd:     true,
		})
	}
	order := &biz.Order{
		UserID:        req.GetUserId(),
		PhoneNumber:   req.GetPhoneNumber(),
		IsPaid:        "waiting",
		StreetAddress: req.GetAddress().GetStreetAddress(),
		City:          req.GetAddress().GetCity(),
		Country:       req.GetAddress().GetCountry(),
		ZipCode:       req.GetAddress().GetZipCode(),
		OrderItems:    orderItemBiz,
	}

	err := s.uc.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	// 告知商品微服务 调整库存
	if state := s.uc.UpdateContent(ctx, updateItems); state {
		return nil, nil
	}

	fmt.Print("OrderInfo Create order", order.OrderId)
	return &pb.PlaceOrderResp{
		OrderId: order.OrderId,
	}, nil
}
func (s *OrderServiceService) ListOrder(ctx context.Context, req *pb.ListOrderReq) (*pb.ListOrderResp, error) {
	var findParam = &biz.FindParams{
		ID:          req.GetUserId(),
		IsPaid:      nil,
		Page:        0,
		PageSize:    0,
		PhoneNumber: "",
	}

	dbOrders, total, err := s.uc.FindOrder(ctx, findParam)
	if err != nil {
		return nil, err
	}
	var orders []*pb.Order
	for _, o := range dbOrders {
		var address = pb.Address{
			City:          o.City,
			Country:       o.Country,
			StreetAddress: o.StreetAddress,
			ZipCode:       o.ZipCode,
		}
		var orderItems []*pb.OrderItem
		for _, oi := range o.OrderItems {
			orderItems = append(orderItems, &pb.OrderItem{
				Cost:      oi.Cost,
				ProductId: oi.ProductId,
				Quantity:  oi.Quantity,
			})
		}
		orders = append(orders, &pb.Order{
			OrderId:     o.OrderId,
			UserId:      o.UserID,
			PhoneNumber: o.PhoneNumber,
			Address:     &address,
			PayState:    o.IsPaid,
			OrderItems:  orderItems,
		})
	}
	return &pb.ListOrderResp{
		Total:  total,
		Orders: orders,
	}, nil
}
func (s *OrderServiceService) MarkOrderPaid(ctx context.Context, req *pb.MarkOrderPaidReq) (*pb.MarkOrderPaidResp, error) {
	var order = &biz.Order{
		OrderId: req.GetOrderId(),
		IsPaid:  "paid",
	}
	err := s.uc.UpdateOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	return &pb.MarkOrderPaidResp{}, nil
}
