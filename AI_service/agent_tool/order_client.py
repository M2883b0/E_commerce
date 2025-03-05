# -*- coding:utf-8 -*-
# @FileName :order_client.py
# @Time :2025/3/4 15:46
# @Author :M2883b0
import grpc
from langchain_core.tools import StructuredTool
from pydantic import BaseModel, Field
from api.order import order_service_pb2_grpc as pb2g
from api.order import order_service_pb2 as pb2
from .etcd_client import get_server_info


# from etcd_client import get_server_info


class OrderItem(BaseModel):
    product_id: int
    quantity: int
    cost: float


class OrderInfo(BaseModel):
    user_id: int
    address: str
    phone_number: str
    order_items: list[OrderItem]


def _create_order(user_id: int, address: str, phone_number: str, order_items: list[OrderItem]):
    # order_server = ""
    order_server = get_server_info("order_service")
    order_items = [dict(item) for item in order_items]
    # for k, v in
    print(f"order_items type is {type(order_items)}")
    if not order_server:
        order_server = "localhost:8060"
    channel = grpc.insecure_channel(order_server)
    stub = pb2g.OrderServiceStub(channel)
    request = pb2.PlaceOrderReq(
        user_id=user_id,
        address={"street_address": address},
        phone_number=phone_number,
        order_items=order_items
    )
    try:
        response = stub.PlaceOrder(request)
    except Exception as e:
        print(e)
        return "failed"
    if response and response.order_id != 0:
        return f"成功创建订单，订单号为 {response.order_id}"
    return f"创建订单失败，请核对订单信息"


create_order = StructuredTool.from_function(
    func=_create_order,
    args_schema=OrderInfo,
    name="create_order",
    description="Create order"
)


class UserId(BaseModel):
    user_id: int


def _list_order(user_id: int):
    order_server = get_server_info("order_service")
    print(f"order_server is {order_server}")
    if not order_server:
        order_server = "localhost:8060"
    channel = grpc.insecure_channel(order_server)
    stub = pb2g.OrderServiceStub(channel)
    request = pb2.ListOrderReq(
        user_id=user_id,
    )
    try:
        response = stub.ListOrder(request)
    except Exception as e:
        print(e)
        return "failed"
    if response and response.total:
        return f"用户 {user_id} 的订单信息为 {response.order_items}"
    return f"用户 {user_id} 没有订单"


list_order = StructuredTool.from_function(
    func=_list_order,
    args_schema=UserId,
    name="list_order",
    description="List order"
)


class OrderId(BaseModel):
    order_id: int


def _get_order_by_id(order_id: int):
    order_server = get_server_info("order_service")
    print(f"order_server is {order_server}")
    if not order_server:
        order_server = "localhost:8060"
    channel = grpc.insecure_channel(order_server)
    stub = pb2g.OrderServiceStub(channel)
    request = pb2.GetOrderByIdReq(
        order_id=order_id,
    )
    try:
        response = stub.GetOrderById(request)
    except Exception as e:
        print(e)
        return "failed"
    assert isinstance(response, pb2.GetOrderByIdResp)
    try:
        if response and response.order and str(response.order):
            return f"订单号为 {order_id} 的订单信息为 {response.order}"
    except Exception as e:
        print(e)
    return f"订单号为 {order_id} 的订单不存在，或者不是您的订单"


get_order_by_id = StructuredTool.from_function(
    func=_get_order_by_id,
    args_schema=OrderId,
    name="get_order_by_id",
    description="Get order by id"
)


def _del_order_by_id(order_id: int):
    order_server = get_server_info("order_service")
    if not order_server:
        order_server = "localhost:8060"
    channel = grpc.insecure_channel(order_server)
    stub = pb2g.OrderServiceStub(channel)
    request = pb2.DelOrderByIdReq(
        order_id=order_id,
    )
    try:
        response = stub.DelOrderById(request)
    except Exception as e:
        print(e)
        return "failed"
    assert isinstance(response, pb2.DelOrderByIdResp)
    if response and response.state != 0:
        return f"成功删除订单，订单号为 {response.order_id}"
    return f"删除订单失败，请核对订单信息"


del_order_by_id = StructuredTool.from_function(
    func=_del_order_by_id,
    args_schema=OrderId,
    name="del_order_by_id",
    description="Del order by id"
)


def _mark_order_cancel(user_id: int, order_id: int):
    order_server = get_server_info("order_service")
    if not order_server:
        order_server = "localhost:8060"
    channel = grpc.insecure_channel(order_server)
    stub = pb2g.OrderServiceStub(channel)
    request = pb2.MarkOrderCancelReq(
        user_id=user_id,
        order_id=order_id,
    )
    try:
        response = stub.MarkOrderCancel(request)
    except Exception as e:
        print(e)
        return "failed"
    assert isinstance(response, pb2.MarkOrderCancelResp)
    if response and response.state != 0:
        return f"成功取消订单，订单号为 {response.order_id}"
    return f"取消订单失败，请核对订单信息"


class MarkOrderReq(BaseModel):
    user_id: int
    order_id: int


mark_order_cancel = StructuredTool.from_function(
    func=_mark_order_cancel,
    args_schema=MarkOrderReq,
    name="mark_order_cancel",
    description="Mark order cancel"
)

if __name__ == '__main__':
    print(_create_order(1234, address="123 Main St", phone_number="322323232332", order_items=[
        {"product_id": 1, "quantity": 2, "cost": 1.1}]))
