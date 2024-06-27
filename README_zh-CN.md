# Catalog Service Management API

* [English](README.md)
* [简体中文](README_zh-CN.md)

https://github.com/daymade/catalog-service-management-api/assets/4291901/f30dd4e7-23d6-4a17-a13d-6c644343b7fd

Catalog-Demo 是一个微服务 API 管理平台，用户可以在前端 Dashboard 管理服务和版本。

本项目是 Catalog-Demo 的后端代码，可以从这里启动整个平台，包括后端，前端和监控。

Demo 包含以下功能：
- Service 的 List 和 Get 接口，支持搜索、过滤、排序、分页、查看详情等功能。
- 基于 API Key 的简单认证机制。
- 支持使用内存和 PostgreSQL 两种存储引擎。
- 测试代码和 swagger 文档。
- Grafana 监控

Demo 中不包含的功能：
- 基于角色的授权机制
- Service 的 CRUD

目前存在的 Bug:
- 前端的 Service 详情页宽度不对，修复起来比较花时间，所以暂时略过，又不是不能用。
- Grafana 可以自动导入数据源，但需要手动导入 Dashboard `build/config/grafana/dashboards/Go Metrics-1719497538877.json`

文件目录：
```
.
├── Makefile # 项目的 makefile 文件，使用 make 命令可以快速运行、测试、构建项目
├── api      # 自动生成的 swagger 文档
├── assets   # 存点图片等静态资源
├── build    # CI/CD 相关，包含 docker file, grafana 和 victoriametrics 的配置文件
├── cmd      # 代码 main 入口
├── docs     # 详细文档
├── internal # 项目大部分代码在这里
├── scripts  # makefile 调用的脚本，包括 docker-compose 和数据库初始化脚本
└── test
```

## Demo 相关背景声明

> 在实际的项目开发中，我们需要就产品细节跟产品经理、设计师、业务运营人员来回沟通，确定产品文档中未能在第一次描述时全部确定的细节，由于项目特殊，我这里简单的假设了一些使用场景，这只是为了减少和面试官中间的沟通损耗。

我们有以下假设：

- 业务定义：

	- 业务我们假设每个 Service 都是一个后端 API 项目，包含了一系列 API 集合
	- 版本管理：Service 有版本管理，版本管理的力度在 Service 级别而不是 API 级别，比如 `/v1` 的 Service 可能包含 10个 API，`/v2` 的Service 可能包含 12 个 API。版本号的规则 `v1` 、`v2`，但可以是任何符合语义化版本的值，我们知道 Google Cloud 的 API 是 `v2024-06-26` 这样。
	- 多租户：只设计最核心的 Service Cards，不需要进行跨区域和多租户设计，比如 Region、Tanent。
	- 权限控制：用户能看到自己的项目，**也可以**看到其他人的项目，实现用户维度的项目过滤不在这一期的考虑范围内。

- 功能需求：

	- 搜索：用户可以通过名字和描述搜索指定 Service
    - 过滤：用户只能通过 Service 的名字和描述进行过滤
    - 排序：用户可以通过名字和创建时间进行排序
    - 分页：由于数据量很小，所以可以支持跳转到指定页，否则只需要支持上一页和下一页
    - 查看详情：用户可以查看 Service 的详情，包括版本、API 列表等
	- 开发者体验：
		- UI: 需要支持 url 规则化，能通过 Uri 跳转到任何中间页面，例如：
			- `services` 是列表页面，如果输入了过滤条件则是 `services?query=name` 。
			- 通过 `services/contact-us` 或者  `services/locate-us` 可以直接跳转到某个 Service 的详情页面。

- 非功能需求：
  - API 规范：我们设计符合 [Google API 规范](https://google.aip.dev/) 的 API。
  - 数据量：
      - 总 Service 数量：10 ～ 10000
      - 总用户数量：1000 以下
      - 每个用户能够创建的 Service 数量有限，最多创建 10 个 service。
      - 每个 Service 的版本数量：最多 10 个版本。

- 技术选型：
	- 搜索：由于数据量很小，我们不引入搜索引擎，直接在数据库上实现过滤, 现阶段也不用考虑索引。
	- 存储引擎：我们支持内存数据库和 PostgreSQL 两种存储引擎，内存数据库用于快速演示，PostgreSQL 可以用于生产环境。
    - 数据库结构：互联网架构中一般不会使用外键，这个数据量很小，外键不会影响性能，所以用了外键。
    - 监控：使用 VictoriaMetrics 和 Grafana 监控服务的性能。

## 运行环境

- Go 1.22 或更高版本
- Docker 和 Docker Compose（使用 PostgreSQL 时需要）, 使用内存数据库时不需要 Docker

## 快速开始

### 运行

1. 以下命令任选其一：

    ```bash
    make run-local # 在本机直接运行 go 代码
    # 或者
	make run-docker # 使用 docker 运行后端和前端
    # 或者
    make run-all # 使用 docker 运行后端、前端和监控，使用内存数据库快速演示
    ```

2. 根据提示选择存储引擎（内存数据库或 PostgreSQL）, 

   1. 如果选择使用内存数据库，除了 go 代码本身以外没有其他依赖，进入第 3 步。

   2. 如果选择 PostgreSQL，脚本将在 Docker 中运行数据库。
      1. 选择是否需要重建数据库，脚本会自动完成建表操作，初次运行不需要选择。
      2. 详细请参考文档：[使用 PostgreSQL 作为存储引擎](docs/postgresql/Use-PostgreSQL.md)

3. 后端 API 将在 `http://localhost:8080` 上可用。
   1. 前端： `http://localhost:5173`
   2. Grafana： `http://localhost:3000`，用户 admin 密码 admin
   3. VictoriaMetrics： `http://localhost:8428`

4. 使用 curl 或 Insomnia 测试端点：

	```bash
	# 测试获取服务列表
	curl -H "Authorization: Bearer dummy_token" http://localhost:8080/api/v1/services
	
	# 测试获取特定服务详情
	curl -H "Authorization: Bearer dummy_token" http://localhost:8080/api/v1/services/1
	```

## 业务建模
```
+-------------------+           +-------------------+
|       User        |           |     Service       |
+-------------------+           +-------------------+
| - id: int         |1         *| - id: int         |
| - name: string    +-----------| - name: string    |
| - email: string   |           | - description: string |
+-------------------+           | - userId: int     |
                                +-------------------+
                                      |1
                                      |
                                      |*
                                +-------------------+
                                |     Version       |
                                +-------------------+
                                | - id: int         |
                                | - version: string |
                                | - serviceId: int  |
                                +-------------------+
                                      |1
                                      |
                                      |*
                                +-------------------+
                                |       API         |
                                +-------------------+
                                | - id: int         |
                                | - name: string    |
                                | - path: string    |
                                | - method: string  |
                                | - versionId: int  |
                                +-------------------+
```
define the domain model of an api management platform,
- include the concept of [user,service,version,api]
- each service can be created by only one user
- each service has multiple version
- each service contains multiple apis, related with specific version

## 架构图

### 类似 [COLA](https://github.com/alibaba/COLA) 的分层架构

以 domain 为核心，在表现层可以有 http api 或 grpc 等不同协议的 adapter。
<img width="566" alt="image" src="https://github.com/daymade/catalog-service-management-api/assets/4291901/4cc9a67b-5356-40a7-840d-6154c8b3d68c">

### 和 Service 相关的类依赖关系

app 层依赖 domain 层的接口，domain 的接口由 infra 层实现，app 负责注入 infra 到 domain，依赖关系为：app -> domain <- infra。
<img width="558" alt="image" src="https://github.com/daymade/catalog-service-management-api/assets/4291901/4e73e449-1e44-4dfa-a957-a5703b1b8ebb">

## API 文档

http://localhost:8080/swagger/index.html

## 开发人员

### 我
<a href="https://github.com/daymade" class="" data-hovercard-type="user" data-hovercard-url="/users/daymade/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self">
  <img src="https://avatars.githubusercontent.com/u/4291901?s=64&amp;v=4" alt="@daymade" width="64" height="64" style="border-radius: 50%; margin-right: 10px;">
</a>

### Claude-3.5-Sonnet
<a href="https://www.anthropic.com/claude" class="" data-hovercard-type="user" data-hovercard-url="/users/claude/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self">
  <img src="https://www.anthropic.com/_next/image?url=https%3A%2F%2Fcdn.sanity.io%2Fimages%2F4zrzovbb%2Fwebsite%2F1c42a8de70b220fc1737f6e95b3c0373637228db-1319x1512.gif&w=3840&q=75" alt="Claude" width="64" height="64" style="border-radius: 50%; margin-right: 10px;">
</a>

### GPT-4o-128k
<a href="https://www.openai.com/gpt-4" class="" data-hovercard-type="user" data-hovercard-url="/users/gpt-4/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self">
  <img src="https://github.com/daymade/catalog-service-management-api/assets/4291901/1bd3390f-4319-44c2-9288-7208e9dc25f8" alt="GPT-4" height="64" style="border-radius: 50%; margin-right: 10px;">
</a>
