FROM golang:1.23 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

COPY --from=builder /src/bin /app
COPY --from=builder /src/configs /data/conf
COPY --from=builder /src/init.sh /

WORKDIR /app

RUN /init.sh

CMD ["./server", "-conf", "/data/conf"]
#CMD ["nc", "-lvp", "9000"]
