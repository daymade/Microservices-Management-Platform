#!/bin/bash

source scripts/base.sh

export COMPOSE_PROJECT_NAME=catalog-service-management-api

setup_storage

# 构建并启动 Docker 容器
log "使用 Docker 构建并启动 app 容器..."
trap 'log "docker-compose 已被用户手动停止。"; exit 0' SIGINT SIGTERM
docker-compose -f build/docker-compose.yml up --build app
