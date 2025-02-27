FROM golang:1.23 AS builder

COPY . /src
WORKDIR /src


RUN mkdir -p bin/ && GOPROXY=https://goproxy.cn go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

FROM debian:stable-slim

COPY --from=builder /src/bin /app

WORKDIR /app

RUN echo -e "deb https://mirrors.aliyun.com/debian/ bookworm main non-free non-free-firmware contrib\n deb-src https://mirrors.aliyun.com/debian/ bookworm main non-free non-free-firmware contrib \n deb https://mirrors.aliyun.com/debian-security/ bookworm-security main \n deb-src https://mirrors.aliyun.com/debian-security/ bookworm-security main \n deb https://mirrors.aliyun.com/debian/ bookworm-updates main non-free non-free-firmware contrib\n deb-src https://mirrors.aliyun.com/debian/ bookworm-updates main non-free non-free-firmware contrib \n deb https://mirrors.aliyun.com/debian/ bookworm-backports main non-free non-free-firmware contrib \n deb-src https://mirrors.aliyun.com/debian/ bookworm-backports main non-free non-free-firmware contrib">/etc/apt/sources.list \
    &&apt-get update && \
    apt-get install -y --no-install-recommends \
    ca-certificates  \
    netbase \
#    netcat-openbsd \
    && rm -rf /var/lib/apt/lists/ \
    && apt-get autoremove -y && apt-get autoclean -y \
    && mv ./* server

CMD ["./server"]
#CMD ["nc", "-lvp", "9000"]