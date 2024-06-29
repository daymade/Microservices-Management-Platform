## `/scripts`

这个目录包含了项目运行和管理所需的各种脚本。

### 环境变量

在脚本中用到了以下环境变量：

- `USE_DB`：设置为 `true` 以使用数据库。
- `DB_HOST`：数据库主机（例如 `localhost` 或 Docker 中的 `db`）。
- `DB_USER`：数据库用户。
- `DB_PASSWORD`：数据库密码。
- `DB_NAME`：数据库名称。
- `DB_PORT`：数据库端口（默认是 `5432`）。
- `COMPOSE_PROJECT_NAME`：Docker Compose 项目名称。

### 文件夹结构和作用

```
scripts/
├── README.md              # 本文档，提供脚本目录的说明
├── base.sh                # 基础脚本，包含其他脚本共用的函数和配置
├── db/                    # 数据库相关脚本和文件
│   ├── migrations/        # 存放数据库迁移文件，用于创建和更新数据库结构。
│   │   ├── 000001_create_tables.down.sql   # 回滚创建表的 SQL
│   │   └── 000001_create_tables.up.sql     # 创建表的 SQL
│   └── testdata/          # 测试数据
│       └── insert_test_data.sql            # 插入测试数据的 SQL
├── run_all.sh             # 启动所有服务，包括数据库、后端 API 和前端 UI。
├── run_docker.sh          # 在 Docker 环境中构建并运行后端 API 服务。
└── run_local.sh           # 在本地环境中运行后端 API 服务。
```
