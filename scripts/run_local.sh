#!/bin/bash

source scripts/base.sh

export COMPOSE_PROJECT_NAME=catalog-service-management-api

setup_storage

# 运行应用程序并处理中断信号
log "启动应用程序..."
trap 'log "应用程序已退出。"; exit 0' SIGINT SIGTERM
# 因为是本机的 go 所以不能解析 db 容器地址，使用 localhost
DB_HOST=localhost go run cmd/server/main.go
