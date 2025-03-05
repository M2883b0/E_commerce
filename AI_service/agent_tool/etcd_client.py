# -*- coding:utf-8 -*-
# @FileName :etcd_client.py
# @Time :2025/3/4 23:00
# @Author :M2883b0
import os

import etcd3
import json


def get_server_info(server_name):
    # return ""
    try:
        client = etcd3.client(host=os.getenv('ETCD_ADDR'), port=2379)
    except Exception as e:
        print(e)
        return ""
    for info, _ in client.get_prefix(f'/microservices/{server_name}'):
        # return json.loads(info.decode())["endpoints"]
        return json.loads(info.decode())["endpoints"][0][7:]
    return ""




def register(server_name, server_addr):
    client = etcd3.client(host=os.getenv('ETCD_ADDR'), port=2379)
    client.put(f'/microservices/{server_name}/3d1d1d1d1d1f', json.dumps({"id":"3d1d1d1d1d1f", "name":server_name, "version":"", "metadata": {}, "endpoints": [server_addr]}))


if __name__ == '__main__':
    print(register("ai_agent", "grpc://127.0.0.1:9000"))
    print(get_server_info("ai_agent"))
