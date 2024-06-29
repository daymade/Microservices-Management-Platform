## `/build`

该目录包含了项目构建、部署和监控相关的配置文件和 Dockerfile。

### 文件夹结构和作用

```
build/
├── Dockerfile                # golang 应用的 Dockerfile
├── README.md                 # 本文档，提供 build 目录的说明
├── config/                   # 配置文件目录
│   ├── grafana/              # Grafana 配置
│   │   ├── dashboards/       # Grafana 仪表板 JSON 文件
│   │   ├── grafana.ini       # Grafana 主配置文件
│   │   └── provisioning/     # Grafana 自动配置
│   ├── otel/                 # OpenTelemetry 配置
│   │   └── otel-collector-config.yaml
│   └── prometheus/           # Prometheus 配置
│       └── prometheus.yml
├── docker-compose.yml        # Docker Compose 配置文件
├── frontend/                 # 前端相关构建文件
│   └── Dockerfile            # 前端应用的 Dockerfile
└── tool/                     # 工具相关构建文件
	└── migrate/              # 数据库迁移工具
		└── Dockerfile        # 迁移工具的 Dockerfile
```

### 文件说明

- `Dockerfile`：用于构建主应用 Docker 镜像的文件。
- `config/`：包含各种服务的配置文件。
  - `grafana/`：Grafana 监控仪表板的配置。
    - `dashboards/`：预定义的 Grafana 仪表板 JSON 文件。
    - `grafana.ini`：Grafana 的主配置文件。
    - `provisioning/`：Grafana 的自动配置文件，用于预加载数据源和仪表板。
  - `otel/`：OpenTelemetry 收集器的配置。
  - `prometheus/`：Prometheus 监控系统的配置。
- `docker-compose.yml`：定义和运行多容器 Docker 应用程序的配置文件。
- `frontend/Dockerfile`：用于构建前端应用 Docker 镜像的文件。
- `tool/migrate/Dockerfile`：用于构建数据库迁移工具 Docker 镜像的文件。

### 使用说明

1. 要构建和运行整个应用程序栈，请使用 `docker-compose`:

   ```shell
    doker-compose up --build
   ```

2. 要单独构建主应用，请在项目根目录运行：

   ```shell
	docker build -f build/Dockerfile -t main-app .
   ```

3. 要构建前端应用，请运行：

   ```shell
	docker build -f build/frontend/Dockerfile -t frontend-app .
   ```

4. 要构建数据库迁移工具，请运行：

   ```shell
	docker build -f build/tool/migrate/Dockerfile -t db-migrate .
   ```

5. Grafana 仪表板可在 `http://localhost:3000` 访问（默认用户名和密码在 `grafana.ini` 中设置）。

6. Prometheus 可在 `http://localhost:9090` 访问。

请确保在运行这些命令之前，已经设置了所有必要的环境变量和配置文件。
