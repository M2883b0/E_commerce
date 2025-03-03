package service

import (
	pb "OrderService/api/order"
	"OrderService/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type OrderServiceService struct {
	pb.UnimplementedOrderServiceServer
	uc *biz.OrderUseCase
}

func NewOrderServiceService(uc *biz.OrderUseCase) *OrderServiceService {
	return &OrderServiceService{uc: uc}
}

func (s *OrderServiceService) PlaceOrder(ctx context.Context, req *pb.PlaceOrderReq) (*pb.PlaceOrderResp, error) {

	var orderItemBiz []*biz.OrderItem           // 商品信息
	var updateItemsSub []*biz.UpdateContentItem // 更新商品信息item
	var updateItemsAdd []*biz.UpdateContentItem // 更新商品信息item
	var checkoutItems []*biz.CheckoutOrderItem  // 结算商品信息item

	for _, orderItemReq := range req.GetOrderItems() {
		orderItemBiz = append(orderItemBiz,
			&biz.OrderItem{
				Quantity:  orderItemReq.GetQuantity(),
				ProductId: orderItemReq.GetProductId(),
				Cost:      orderItemReq.GetCost(),
			})
		updateItemsSub = append(updateItemsSub, &biz.UpdateContentItem{
			ProductId: orderItemReq.GetProductId(),
			Quantity:  int32(orderItemReq.GetQuantity()),
			IsAdd:     false,
		})
		updateItemsAdd = append(updateItemsAdd, &biz.UpdateContentItem{
			ProductId: orderItemReq.GetProductId(),
			Quantity:  int32(orderItemReq.GetQuantity()),
			IsAdd:     true,
		})
		checkoutItems = append(checkoutItems, &biz.CheckoutOrderItem{
			ProductId: orderItemReq.GetProductId(),
			Price:     orderItemReq.GetCost(),
			Quantity:  int32(orderItemReq.GetQuantity()),
		})
	}

	// 告知商品微服务 调整库存
	if !s.uc.UpdateContent(ctx, updateItemsSub) {
		log.Infof("调整库存失败")
		return nil, nil
	}

	// 结算，如果结算失败，需要再次告诉商品微服务，库存回滚
	checkoutRsp, err := s.uc.CheckoutOrder(ctx, checkoutItems)
	if err != nil || checkoutRsp.HasChanged {
		if !s.uc.UpdateContent(ctx, updateItemsAdd) {
			log.Infof("回滚库存失败")
		}
		// TODO 记得注释回来
		//return nil, nil
	}

	// 数据库创建订单
	order := &biz.Order{
		UserID:         req.GetUserId(),
		PhoneNumber:    req.GetPhoneNumber(),
		ActualPayment:  checkoutRsp.ActualPrice,
		OriginalCharge: checkoutRsp.TotalPrice,
		IsFreeShipping: checkoutRsp.IsFreeShipping,
		ShippingFee:    checkoutRsp.ShippingFee,
		OrderState:     "waiting",
		StreetAddress:  req.GetAddress().GetStreetAddress(),
		City:           req.GetAddress().GetCity(),
		Country:        req.GetAddress().GetCountry(),
		ZipCode:        req.GetAddress().GetZipCode(),
		OrderItems:     orderItemBiz,
	}

	err = s.uc.CreateOrder(ctx, order)
	if err != nil {
		log.Infof("数据库创建订单失败")
		return &pb.PlaceOrderResp{
			OrderId: 0,
		}, nil
	}

	log.Infof("成功创建订单 订单号为 %+v", order.OrderId)
	return &pb.PlaceOrderResp{
		OrderId: order.OrderId,
	}, nil
}
func (s *OrderServiceService) ListOrder(ctx context.Context, req *pb.ListOrderReq) (*pb.ListOrderResp, error) {
	dbOrders, total := s.uc.FindOrderByUserId(ctx, req.GetUserId())
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
			OrderId:        o.OrderId,
			UserId:         o.UserID,
			OriginalCharge: o.OriginalCharge,
			ActualPayment:  o.ActualPayment,
			IsFreeShipping: o.IsFreeShipping,
			ShippingFee:    o.ShippingFee,
			PhoneNumber:    o.PhoneNumber,
			Address:        &address,
			OrderState:     o.OrderState,
			OrderItems:     orderItems,
		})
	}
	return &pb.ListOrderResp{
		Total:  total,
		Orders: orders,
	}, nil
}

func (s *OrderServiceService) GetOrderById(ctx context.Context, req *pb.GetOrderByIdReq) (*pb.GetOrderByIdResp, error) {
	dbOrders := s.uc.FindOrderById(ctx, req.GetUserId(), req.GetOrderId())
	var address = pb.Address{
		City:          dbOrders.City,
		Country:       dbOrders.Country,
		StreetAddress: dbOrders.StreetAddress,
		ZipCode:       dbOrders.ZipCode,
	}
	var orderItems []*pb.OrderItem
	for _, oi := range dbOrders.OrderItems {
		orderItems = append(orderItems, &pb.OrderItem{
			Cost:      oi.Cost,
			ProductId: oi.ProductId,
			Quantity:  oi.Quantity,
		})

	}
	order := pb.Order{
		OrderId:        dbOrders.OrderId,
		UserId:         dbOrders.UserID,
		OriginalCharge: dbOrders.OriginalCharge,
		ActualPayment:  dbOrders.ActualPayment,
		IsFreeShipping: dbOrders.IsFreeShipping,
		ShippingFee:    dbOrders.ShippingFee,
		PhoneNumber:    dbOrders.PhoneNumber,
		Address:        &address,
		OrderState:     dbOrders.OrderState,
		OrderItems:     orderItems,
	}

	return &pb.GetOrderByIdResp{
		Order: &order,
	}, nil
}

func (s *OrderServiceService) DelOrderById(ctx context.Context, req *pb.DelOrderByIdReq) (*pb.DelOrderByIdResp, error) {
	dbOrder := s.uc.FindOrderById(ctx, req.GetUserId(), req.GetOrderId())
	if dbOrder.OrderId == -1 {
		return &pb.DelOrderByIdResp{
			State: false,
		}, nil
	}

	err := s.uc.DeleteOrder(ctx, req.GetOrderId())
	if err != nil {
		return &pb.DelOrderByIdResp{
			State: false,
		}, nil
	}

	return &pb.DelOrderByIdResp{
		State: true,
	}, nil
}
func (s *OrderServiceService) MarkOrderPaid(ctx context.Context, req *pb.MarkOrderPaidReq) (*pb.MarkOrderPaidResp, error) {
	log.Infof("mark order paid. order id is %+v", req.OrderId)
	var order = &biz.Order{
		OrderId:    req.GetOrderId(),
		OrderState: "paid",
	}
	err := s.uc.UpdateOrder(ctx, order)
	if err != nil {
		return &pb.MarkOrderPaidResp{
			State: false,
		}, nil
	}
	return &pb.MarkOrderPaidResp{
		State: true,
	}, nil
}

func (s *OrderServiceService) MarkOrderCancel(ctx context.Context, req *pb.MarkOrderCancelReq) (*pb.MarkOrderCancelResp, error) {
	log.Infof("mark order cancel. order id is %+v, user id is %+v", req.OrderId, req.OrderId)
	var order = &biz.Order{
		OrderId:    req.GetOrderId(),
		OrderState: "cancel",
	}
	err := s.uc.UpdateOrder(ctx, order)
	if err != nil {
		return &pb.MarkOrderCancelResp{
			State: false,
		}, nil
	}
	return &pb.MarkOrderCancelResp{
		State: true,
	}, nil
}
