# 使用 PostgreSQL 作为存储引擎

> ![NOTICE]
> 本文档适合需要了解数据库迁移的详细信息，没有出问题可以忽略本文档。

如果选择使用 PostgreSQL 作为存储引擎，项目会自动创建数据库和相关表。

## 数据库迁移

本项目使用 `golang-migrate` 管理数据库迁移。

## 安装

在 macOS 上安装 `golang-migrate`：

```bash
brew install golang-migrate
```

## 应用迁移

在使用 PostgreSQL 作为存储引擎时，您可以选择应用所有未执行的迁移：

```bash
cd catalog-service-management-api
migrate -path init/db/migrations -database postgres://user:password@localhost:5432/services_db?sslmode=disable up
```

## 回滚迁移

回滚最近一次的迁移：

```bash
migrate -path init/db/migrations -database postgres://user:password@localhost:5432/services_db?sslmode=disable down
```

## 查看迁移状态

检查迁移状态：

```bash
migrate -path init/db/migrations -database "postgres://user:password@localhost:5432/services_db?sslmode=disable" version
```

有关更多 `golang-migrate` 的命令和选项，请参考 [官方文档](https://github.com/golang-migrate/migrate)。
