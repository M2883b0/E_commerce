#!/bin/bash

# 备份原有的 sources.list 文件
cp /etc/apt/sources.list /etc/apt/sources.list.bak

# 使用阿里云的 Debian 镜像源
cat > /etc/apt/sources.list <<EOF
deb http://mirrors.163.com/debian/ bookworm main non-free non-free-firmware contrib
deb-src http://mirrors.163.com/debian/ bookworm main non-free non-free-firmware contrib
deb http://mirrors.163.com/debian-security/ bookworm-security main
deb-src http://mirrors.163.com/debian-security/ bookworm-security main
deb http://mirrors.163.com/debian/ bookworm-updates main non-free non-free-firmware contrib
deb-src http://mirrors.163.com/debian/ bookworm-updates main non-free non-free-firmware contrib
deb http://mirrors.163.com/debian/ bookworm-backports main non-free non-free-firmware contrib
deb-src http://mirrors.163.com/debian/ bookworm-backports main non-free non-free-firmware contrib
EOF

# 更新软件包列表
apt-get update

apt-get install -y \
ca-certificates  \
    netbase \
    netcat-openbsd \

rm -rf /var/lib/apt/lists/
apt-get autoremove -y
apt-get autoclean -y

mv /app/* /app/server
