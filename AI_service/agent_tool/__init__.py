# -*- coding:utf-8 -*-
# @FileName :__init__.py.py
# @Time :2025/3/4 15:41
# @Author :M2883b0
from .order_client import *
from .product_client import *
from .user_client import *

tools = [
    create_order,
    del_order_by_id,
    list_order,
    mark_order_cancel,
    get_order_by_id,
find_user_info_by_id,
    query_product_by_id,
    find_product_by_decription
]

if __name__ == '__main__':
    pass
