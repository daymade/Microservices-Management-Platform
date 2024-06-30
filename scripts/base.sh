#!/bin/bash

set -e

log() {
  echo "[`date +"%Y-%m-%d %H:%M:%S"`] $1"
}

setup_db() {
  log "使用 Docker 构建并启动 postgres 容器（常驻）..."
  docker-compose -f build/docker-compose.yml up -d db

  # 设置环境变量
  export USE_DB=true
  export DB_HOST=db
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
}

perform_migration() {
  log "开始清理目标表..."
  docker exec -i catalog-service-management-api-db-1 psql -U $DB_USER -d $DB_NAME -c "DROP TABLE IF EXISTS versions, services, users, schema_migrations CASCADE;" && log "目标表已删除。" || { log "清理表失败"; exit 1; }

  log "执行数据库迁移..."
  docker-compose -f build/docker-compose.yml run --rm migrate -path /migrations -database "postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable" up || { log "数据库迁移失败"; exit 1; }
  log "数据库迁移完成。"

  log "验证迁移是否成功..."
  docker exec -i catalog-service-management-api-db-1 psql -U $DB_USER -d $DB_NAME -c "\dt" | grep -q users
  if [ $? -eq 0 ]; then
    log "迁移成功，插入测试数据..."
    docker exec -i catalog-service-management-api-db-1 psql -U $DB_USER -d $DB_NAME < scripts/db/testdata/insert_test_data.sql
    log "测试数据插入完成。"
  else
    log "迁移失败，找不到表 'users'。"
    exit 1
  fi
}

setup_storage() {
  read -p "请选择存储引擎 (m: memory, p: postgres) [默认 m: memory]: " storage_engine
  storage_engine=${storage_engine:-m}

  if [[ "$storage_engine" == "postgres" || "$storage_engine" == "p" ]]; then
    setup_db

    read -p "是否重建数据库？如果重建会丢失所有数据，初次运行请选 yes (yes/no, 默认: no) [y/N]: " perform_migration
    perform_migration=${perform_migration:-n}

    if [[ "$perform_migration" == "y" || "$perform_migration" == "Y" || "$perform_migration" == "yes" || "$perform_migration" == "YES" ]]; then
      perform_migration
    else
      log "跳过数据库迁移。"
    fi
  else
    log "使用内存数据库"
    export USE_DB=false
  fi
}
