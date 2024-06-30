#!/bin/bash
set -e

# 设置环境变量
export COMPOSE_PROJECT_NAME=catalog-service-management-api-test
export USE_DB=true
export DB_HOST=localhost
export DB_USER=test
export DB_PASSWORD=test
export DB_NAME=testdb
export DB_PORT=5433

# 启动测试数据库
docker-compose -f build/test/docker-compose.test.yml up -d postgres_test

# 等待数据库启动
echo "Waiting for database to start..."
until docker exec ${COMPOSE_PROJECT_NAME}-postgres_test-1 pg_isready -U $DB_USER -d $DB_NAME; do
  sleep 2
done
echo "Database is ready."

# 运行数据库迁移
docker-compose -f build/test/docker-compose.test.yml run --rm migrate \
  -path /migrations \
  -database "postgres://$DB_USER:$DB_PASSWORD@postgres_test:5432/$DB_NAME?sslmode=disable" \
  up

# 插入测试数据
docker exec -i ${COMPOSE_PROJECT_NAME}-postgres_test-1 psql -U $DB_USER -d $DB_NAME < scripts/db/testdata/insert_test_data.sql

echo "Test environment is ready."
