# E_commerce
go语言开发的抖音电商


技术选型
- 微服务：Kratos、etcd
- 数据库：Gorm、MySQL
- Web框架：Gin
- 搜索引擎：Elasticsearch、Kibana
- 缓存：Redis
- 中间件：JWT、Zipkin
- 监控与可视化：Prometheus、Grafana
- AI Agent 框架： LangChain
- 容器化部署：Docker

![](.\技术选型.png)

服务之间的逻辑关系：

![](./逻辑图.png)





一键部署

```bash
git clone https://github.com/M2883b0/E_commerce
# 配置大模型参数
cat > E_commerce/AI_service/.env <<EOF
ARK_API_KEY=33820f8b-9106-4e5d-9f98-c39739ca304a # api key
ARK_BASE_URL=https://ark.cn-beijing.volces.com/api/v3 # 模型地址
MODEL_NAME=doubao-1-5-pro-32k-250115 # 模型名称
EOF
# 一键部署脚本
./one-click_deployment.sh
```