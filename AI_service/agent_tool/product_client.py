# -*- coding:utf-8 -*-
# @FileName :product_client.py
# @Time :2025/3/4 15:50
# @Author :M2883b0
import grpc
from langchain_core.tools import StructuredTool
from pydantic import BaseModel
from api.operate import app_pb2 as pb2, app_pb2_grpc as pb2g
from agent_tool import get_server_info


class FindProductInfo(BaseModel):
    product_description: str


def _find_product_by_decription(product_description: str):
    product_server = get_server_info("content_manage")
    print(f"product_server is {product_server}")
    if not product_server:
        product_server = "localhost:8081"
    channel = grpc.insecure_channel(product_server)
    stub = pb2g.AppStub(channel)
    request = pb2.FindContentReq(
        query=product_description,
    )
    try:
        response = stub.FindContent(request)
    except Exception as e:
        print(e)
        return "失败，服务内部出错"
    try:
        if response and response.total:
            return f"符合描述的商品如下 {response.contents}"
    except Exception as e:
        print(e)
    return f"没有符合商品描述 {product_description} 的商品"


find_product_by_decription = StructuredTool.from_function(
        func=_find_product_by_decription,
        args_schema=FindProductInfo,
        name="find_product_by_decription",
        description="使用商品描述查找商品，会返回{'商品名':id}字典格式，如果用户需要下订单则使用第一个元素即可"
    )


class ProductId(BaseModel):
    product_id: int


def _query_product_by_id(product_id: int):
    product_server = get_server_info("content_manage")
    if not product_server:
        product_server = "localhost:8081"
    channel = grpc.insecure_channel(product_server)
    stub = pb2g.AppStub(channel)
    request = pb2.GetContentReq(
        id=[product_id],
    )
    try:
        response = stub.GetContent(request)
        if response and response.contents:
            return response.contents[0]
    except Exception as e:
        print(e)
        return "失败，服务内部出错"
    return {"商品价格为": 79000}


query_product_by_id = StructuredTool.from_function(
        func=_query_product_by_id,
        args_schema=ProductId,
        name="query_product_by_id",
        description="Query product by id"
    )


if __name__ == '__main__':
    pass
