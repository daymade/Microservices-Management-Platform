# 使用多阶段构建优化镜像大小和构建速度
FROM golang:1.22.2-alpine AS builder

# 根据 CPU 架构下载适合的 migrate 工具
RUN ARCH=$(uname -m) && \
    if [ "$ARCH" = "x86_64" ]; then \
        MIGRATE_URL="https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz"; \
    elif [ "$ARCH" = "aarch64" ]; then \
        MIGRATE_URL="https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-arm64.tar.gz"; \
    else \
        echo "unsupported architecture: $ARCH"; exit 1; \
    fi && \
    wget -O migrate.tar.gz $MIGRATE_URL && \
    tar -xvf migrate.tar.gz -C /usr/local/bin && \
    chmod +x /usr/local/bin/migrate && \
    rm migrate.tar.gz

# 设置工作目录
WORKDIR /app

# 仅复制 go.mod 和 go.sum 并下载依赖，这样可以利用缓存
COPY ../go.mod go.sum ./
RUN go mod download

# 复制其余文件并构建应用
COPY .. .
RUN go build -o main ./cmd/server/main.go

# 使用一个更小的基础镜像来运行构建的应用
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /usr/local/bin/migrate .

EXPOSE 8080

CMD ["./main"]
