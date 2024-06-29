#!/bin/bash

source scripts/base.sh

export COMPOSE_PROJECT_NAME=catalog-service-management-api

setup_storage

# 检查并克隆前端项目
FRONTEND_DIR="build/frontend/catalog-service-management-ui"
if [ -d "$FRONTEND_DIR" ]; then
  log "前端项目目录已存在，更新项目..."
  cd $FRONTEND_DIR
  git pull
else
  log "前端项目目录不存在，克隆项目..."
  git clone https://github.com/daymade/catalog-service-management-ui.git $FRONTEND_DIR
  cd $FRONTEND_DIR
fi
cd -

# 启动 Docker Compose
log "使用 Docker Compose 启动所有服务..."
trap 'log "docker-compose 已被用户手动停止。"; exit 0' SIGINT SIGTERM
docker-compose -f build/docker-compose.yml up --build
