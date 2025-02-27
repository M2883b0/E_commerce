FROM golang:1.23 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

COPY --from=builder /src/bin /app
COPY --from=builder /src/configs /data/conf

WORKDIR /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    ca-certificates  \
    netbase \
#    netcat-openbsd \
    && rm -rf /var/lib/apt/lists/ \
    && apt-get autoremove -y && apt-get autoclean -y \
    && mv ./* server

CMD ["./server", "-conf", "/data/conf"]
#CMD ["nc", "-lvp", "9000"]
