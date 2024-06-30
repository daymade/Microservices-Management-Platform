#!/bin/bash

# 设置错误处理
set -e

# 定义清理函数
cleanup() {
    echo "Tearing down test environment..."
    ./scripts/test/integration/teardown_test_env.sh
}

# 设置 trap，无论脚本如何退出都会执行清理函数
trap cleanup EXIT

echo "Setting up test environment..."
./scripts/test/integration/setup_test_env.sh

echo "Running integration tests..."
USE_DB=true DB_HOST=localhost DB_USER=test DB_PASSWORD=test DB_NAME=testdb DB_PORT=5433 \
    go test -v -tags=integration \
      ./test/integration/... \
      ./internal/infrastructure/storage/...

# 不需要在这里调用清理函数，trap 会确保它被执行
