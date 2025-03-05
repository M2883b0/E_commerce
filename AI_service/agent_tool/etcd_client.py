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
    for info, _ in client.get_prefix(f'/microservices/user_manage'):
        # return json.loads(info.decode())["endpoints"]
        return json.loads(info.decode())["endpoints"][0]
    return ""


def register(server_name, server_addr):
    client = etcd3.client(host=os.getenv('ETCD_ADDR'), port=2379)
    client.put(f'/microservices/{server_name}', json.dumps({"endpoints": [server_addr]}))


if __name__ == '__main__':
    print(register("ai_agent", "127.0.0.1:9000"))
    print(get_server_info("content_manage"))
