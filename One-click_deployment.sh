#!/bin/bash

cmd="docker-compose build --no-cache && docker-compose up"

folders=()
for dir in */; do
    folders+=("${dir%/}")
done

for folder in "${folders[@]}"; do
    screen -dmS "$folder" bash -c "cd \"$folder\" && $cmd"
done

