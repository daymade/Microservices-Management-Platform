#!/bin/bash

set -e

log() {
  echo "[`date +"%Y-%m-%d %H:%M:%S"`] $1"
}

export COMPOSE_PROJECT_NAME=catalog-service-management-api

# 选择存储引擎，提供默认值 memory
read -p "请选择存储引擎 (m: memory, p: postgres) [默认 m: memory]: " storage_engine
storage_engine=${storage_engine:-m}

if [[ "$storage_engine" == "postgres" || "$storage_engine" == "p" ]]; then
  log "使用 Docker 构建并启动 postgres 容器（常驻）..."
  docker-compose -f build/docker-compose.yml up -d db

  # 设置环境变量
  export USE_DB=true
  export DB_HOST=localhost
  export DB_USER=user
  export DB_PASSWORD=password
  export DB_NAME=services_db
  export DB_PORT=5432

  # 等待数据库启动
  log "等待数据库启动..."
  until docker exec catalog-service-management-api-db-1 pg_isready -U $DB_USER -d $DB_NAME; do
    sleep 2
  done
  log "数据库已启动。"

  # 询问用户是否执行数据库迁移，提供默认值 n
  read -p "是否重建数据库？如果重建会丢失所有数据，初次运行请选 yes (yes/no, 默认: no) [y/N]: " perform_migration
  perform_migration=${perform_migration:-n}

  if [[ "$perform_migration" == "y" || "$perform_migration" == "Y" || "$perform_migration" == "yes" || "$perform_migration" == "YES" ]]; then
    log "开始清理目标表..."
    docker exec -i catalog-service-management-api-db-1 psql -U $DB_USER -d $DB_NAME -c "DROP TABLE IF EXISTS versions CASCADE;" && log "表 versions 已删除。" || { log "清理表 versions 失败"; exit 1; }
    docker exec -i catalog-service-management-api-db-1 psql -U $DB_USER -d $DB_NAME -c "DROP TABLE IF EXISTS services CASCADE;" && log "表 services 已删除。" || { log "清理表 services 失败"; exit 1; }
    docker exec -i catalog-service-management-api-db-1 psql -U $DB_USER -d $DB_NAME -c "DROP TABLE IF EXISTS users CASCADE;" && log "表 users 已删除。" || { log "清理表 users 失败"; exit 1; }
    docker exec -i catalog-service-management-api-db-1 psql -U $DB_USER -d $DB_NAME -c "DROP TABLE IF EXISTS schema_migrations CASCADE;" && log "表 schema_migrations 已删除。" || { log "清理表 schema_migrations 失败"; exit 1; }
    log "目标表清理完成。"

    log "执行数据库迁移..."
    migrate -path scripts/db/migrations -database "postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable" up || { log "数据库迁移失败"; exit 1; }
    log "数据库迁移完成。"

    log "验证迁移是否成功..."
    docker exec -i catalog-service-management-api-db-1 psql -U $DB_USER -d $DB_NAME -c "\dt" | grep -q users
    if [ $? -eq 0 ]; then
      log "迁移成功，插入测试数据..."

      # 插入测试数据
      docker exec -i catalog-service-management-api-db-1 psql -U $DB_USER -d $DB_NAME <<EOSQL
INSERT INTO users (username, email) VALUES
('testuser1', 'testuser1@example.com'),
('testuser2', 'testuser2@example.com');
INSERT INTO services (name, description, owner_id) VALUES
('testservice1', 'This is a test service 1', (SELECT id FROM users WHERE username = 'testuser1')),
('testservice2', 'This is a test service 2', (SELECT id FROM users WHERE username = 'testuser2'));
INSERT INTO versions (service_id, number, description) VALUES
((SELECT id FROM services WHERE name = 'testservice1'), 'v1.0', 'Initial version of test service 1'),
((SELECT id FROM services WHERE name = 'testservice2'), 'v1.0', 'Initial version of test service 2');
EOSQL
      log "测试数据插入完成。"
    else
      log "迁移失败，找不到表 'users'。"
      exit 1
    fi
  else
    log "跳过数据库迁移。"
  fi
else
  # 使用内存数据库
  log "使用内存数据库"
  export USE_DB=false
fi

# 运行应用程序并处理中断信号
log "启动应用程序..."
trap 'log "应用程序已退出。"; exit 0' SIGINT SIGTERM
go run cmd/server/main.go
