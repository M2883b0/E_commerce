# -*- coding:utf-8 -*-
# @FileName :user_client.py
# @Time :2025/3/5 1:23
# @Author :M2883b0
import logging

import grpc
from langchain_core.tools import StructuredTool
from pydantic import BaseModel
from api.operate import user_pb2_grpc as pb2g
from api.operate import user_pb2 as pb2
from agent_tool import get_server_info


def _find_user_info_by_id(user_id: int):
    user_server = get_server_info("user_manage")
    if not user_server:
        user_server = "127.0.0.1:8061"
    channel = grpc.insecure_channel(user_server)
    stub = pb2g.UserStub(channel)
    request = pb2.GetUserRequest(
        id=user_id,
    )
    logging.info(f"用户微服务地址为{user_server}")
    print(f"用户微服务地址为{user_server}")
    try:
        response = stub.GetUser(request)
    except Exception as e:
        print(e)
        return "failed"
    assert isinstance(response, pb2.GetUserReply)
    if response and response.user and str(response.user):
        return f"用户 {user_id} 的信息为 {response.user}"
    return f"用户 {user_id} 不存在"


class UserId(BaseModel):
    user_id: int


find_user_info_by_id = StructuredTool.from_function(

    description="根据用户id查找用户信息",
    func=_find_user_info_by_id,
    args_schema=UserId,
    name="find_user_info_by_id",
    return_name="user_info"
)

if __name__ == '__main__':
    pass
