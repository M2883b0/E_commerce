# E_commerce
go语言开发的抖音电商



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