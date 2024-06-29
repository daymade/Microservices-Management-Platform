# 依赖安装阶段
FROM golang:1.22.2-alpine AS deps
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# 构建阶段
FROM golang:1.22.2-alpine AS builder
WORKDIR /app
COPY --from=deps /go/pkg /go/pkg
# 只复制必要的源代码文件和文件夹
COPY api/ ./api/
COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY go.mod go.sum ./
# 如果 assets 文件夹包含构建所需的资源，也复制它
COPY assets/ ./assets/
RUN go build -o main ./cmd/server/main.go

# 最终镜像
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
# 如果需要在运行时使用 assets，也复制它们
COPY --from=builder /app/assets ./assets
EXPOSE 8080
CMD ["./main"]
