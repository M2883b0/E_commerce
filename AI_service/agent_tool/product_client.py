# -*- coding:utf-8 -*-
# @FileName :product_client.py
# @Time :2025/3/4 15:50
# @Author :M2883b0
from langchain_core.tools import StructuredTool
from pydantic import BaseModel


class FindProductInfo(BaseModel):
    product_description: str


def _find_product_by_decription(product_description: str):
    return {"苹果": 79000}


find_product_by_decription = StructuredTool.from_function(
        func=_find_product_by_decription,
        args_schema=FindProductInfo,
        name="find_product_by_decription",
        description="使用商品描述查找商品，会返回{'商品名':id}字典格式，提取第一个元素即可"
    )


class ProductId(BaseModel):
    product_id: int


def _query_product_by_id(product_id: int):
    return {"商品价格为": 79000}


query_product_by_id = StructuredTool.from_function(
        func=_query_product_by_id,
        args_schema=ProductId,
        name="query_product_by_id",
        description="Query product by id"
    )


if __name__ == '__main__':
    pass
