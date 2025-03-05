# -*- coding:utf-8 -*-
# @FileName :app.py
# @Time :2025/3/4 15:25
# @Author :M2883b0
import os

if not os.getenv("ETCD_ADDR"):
    import dotenv

    dotenv.load_dotenv()

import logging

import grpc
from concurrent import futures
from api.ai_agent import ai_agent_pb2_grpc
from api.ai_agent import ai_agent_pb2
from AI_agent import Agent
from agent_tool import etcd_client

logger = logging.Logger("grpc_server", level=logging.DEBUG)
handler = logging.StreamHandler()
formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
handler.setFormatter(formatter)
logger.addHandler(handler)


class GreeterServicer(ai_agent_pb2_grpc.AiAgentServicer):
    def UserRequest(self, request, context):
        logger.info(f"收到用户请求 {request.user_id}")
        agent = Agent()
        response = agent.run(request.user_message, request.user_id)["output"]
        print(f"response is {response}")
        return ai_agent_pb2.UserRequestResp(
            agent_response=response
        )
        # return ai_agent_pb2.UserRequestResp(
        #     agent_response="yes"
        # )

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    logger.info("正在读取配置")
    ai_agent_pb2_grpc.add_AiAgentServicer_to_server(GreeterServicer(), server)
    etcd_client.register("ai_agent", os.getenv("SERVER_ADDR"))
    server.add_insecure_port('[::]:9000')  # 绑定端口
    logger.info("服务器正在监听 9000 端口")
    server.start()
    logger.info("服务器已启动")
    server.wait_for_termination()  # 阻塞直至终止


if __name__ == '__main__':
    serve()
