#!/bin/bash

git fetch origin
git reset --hard origin/main

cmd="docker-compose build --no-cache && docker-compose up"
folders=()

# 遍历文件夹并过滤非目录项
for dir in */; do
    [[ -d "$dir" ]] && folders+=("${dir%/}")
done

for folder in "${folders[@]}"; do
    echo "启动会话: $folder"
    screen -dmS "$folder" bash -c "cd \"$folder\" && $cmd; exec bash"
done

docker-compose up -d

