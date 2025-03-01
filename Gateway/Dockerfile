FROM golang:1.23 AS builder

COPY . /src
WORKDIR /src


RUN mkdir -p bin/ && GOPROXY=https://goproxy.cn go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

FROM debian:stable-slim

COPY --from=builder /src/bin /app
COPY --chmod=755 --from=builder /src/init.sh /

WORKDIR /app

RUN chmod +x /init.sh && /init.sh

CMD ["./server"]
#CMD ["nc", "-lvp", "9000"]
