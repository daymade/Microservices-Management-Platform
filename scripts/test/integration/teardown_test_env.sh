#!/bin/bash

set -e

export COMPOSE_PROJECT_NAME=catalog-service-management-api-test

# 停止并删除测试容器和相关卷
docker-compose -f build/test/docker-compose.test.yml down -v

# 清理可能遗留的卷
volumes=$(docker volume ls -q -f name=${COMPOSE_PROJECT_NAME})
if [ -n "$volumes" ]; then
    echo "Removing leftover volumes..."
    docker volume rm $volumes
fi

# 清理可能遗留的网络
networks=$(docker network ls -q -f name=${COMPOSE_PROJECT_NAME})
if [ -n "$networks" ]; then
    echo "Removing leftover networks..."
    docker network rm $networks
fi

echo "Test environment has been completely torn down."
