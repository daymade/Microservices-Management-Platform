# Catalog Service Management API

* [English](README.md)
* [简体中文](README_zh.md)

本项目是一个微服务 API 管理平台的后端实现，用户可以在前端 Dashboard 管理服务和版本。

项目目前包含以下功能：
- Service 的 List 和 Get 接口，支持搜索、过滤、排序、分页、查看详情等功能。
- 基于 API Key 的简单认证机制，基于角色的用户授权机制。
- 支持使用内存和 PostgreSQL 两种存储引擎。
- 包含测试代码
- 支持 swagger 文档

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

	- 搜索：由于数据量很小，我们不引入搜索引擎，直接在数据库上实现过滤。



## 运行环境

- Go 1.22 或更高版本
- Docker 和 Docker Compose（使用 PostgreSQL 时需要）, 使用内存数据库时不需要 Docker

## 快速开始

### 运行

1. 以下命令任选其一：

    ```bash
    make run-local # 在本机直接运行 go 代码
    # 或者
	make run-docker # 使用 docker 运行
    ```

2. 根据提示选择存储引擎（内存数据库或 PostgreSQL）, 

   1. 如果选择使用内存数据库，除了 go 代码本身以外没有其他依赖，进入第 3 步。

   2. 如果选择 PostgreSQL，脚本将在 Docker 中运行数据库。
      1. 选择是否需要重建数据库，脚本会自动完成建表操作，初次运行不需要选择。
      2. ，并插入一些测试数据。
      2. 详细请参考文档：使用 PostgreSQL 作为存储引擎

3. 应用程序将在 `http://localhost:8080` 上可用。

4. 使用 curl 或 Insomnia 测试端点：

	```bash
	# 测试获取服务列表
	curl -H "Authorization: Bearer dummy_token" http://localhost:8080/api/v1/services
	
	# 测试获取特定服务详情
	curl -H "Authorization: Bearer dummy_token" http://localhost:8080/api/v1/services/1
	```

## 业务建模

define the domain model of an api management platform,
- include the concept of [user,service,version,api]
- each service can be created by only one user
- each service has multiple version
- each service contains multiple apis, related with specific version

## 架构图


## API 文档

http://localhost:8080/swagger/index.html

## 开发

make 命令：
1. build production 会设置变量 `export GIN_MODE=release`
