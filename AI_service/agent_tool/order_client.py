# -*- coding:utf-8 -*-
# @FileName :order_client.py
# @Time :2025/3/4 15:46
# @Author :M2883b0
from langchain_core.tools import StructuredTool
from pydantic import BaseModel


#  创建订单工具
class Address(BaseModel):
    street_address: str
    city: str
    country: str
    zip_code: int


class OrderItem(BaseModel):
    product_id: int
    quantity: int
    address: str


class OrderInfo(BaseModel):
    user_id: int
    address: Address
    phone_number: str
    order_items: list[OrderItem]


def _create_order(user_id: int, address: Address, phone_number: str, order_items: list[OrderItem]):
    return "success"


create_order = StructuredTool.from_function(
    func=_create_order,
    args_schema=OrderInfo,
    name="create_order",
    description="Create order"
)

if __name__ == '__main__':
    pass
