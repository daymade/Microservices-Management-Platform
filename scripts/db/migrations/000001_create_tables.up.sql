CREATE TABLE users (
                       id BIGSERIAL PRIMARY KEY,
                       username VARCHAR(100) NOT NULL UNIQUE,
                       email VARCHAR(100) NOT NULL UNIQUE,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE services (
                          id BIGSERIAL PRIMARY KEY,
                          name VARCHAR(100) NOT NULL,
                          description TEXT,
                          owner_id BIGINT REFERENCES users(id),
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE versions (
                          id BIGSERIAL PRIMARY KEY,
                          service_id BIGINT REFERENCES services(id),
                          number VARCHAR(50) NOT NULL,
                          description TEXT,
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 为了优化模糊搜索操作，需要在 PostgreSQL 中创建 `pg_trgm` 扩展和索引。
-- 这行命令安装 pg_trgm 扩展，它提供了 trigram 匹配功能，用于优化模糊搜索。
-- pg_trgm 扩展通常已经包含在官方的 PostgreSQL Docker 镜像（如 postgres:13、postgres:14 等）中，只是默认没有启用。
CREATE EXTENSION IF NOT EXISTS pg_trgm;

-- 为 services 表的 name 和 description 列添加 GIN (Generalized Inverted Index) 索引，使用 trigram 操作符 gin_trgm_ops。
CREATE INDEX IF NOT EXISTS idx_services_name_trgm ON services USING gin (name gin_trgm_ops);
CREATE INDEX IF NOT EXISTS idx_services_description_trgm ON services USING gin (description gin_trgm_ops);

-- 普通的 B-tree 索引，用于优化基于 created_at 的排序和范围查询。
CREATE INDEX IF NOT EXISTS idx_services_created_at ON services (created_at);
