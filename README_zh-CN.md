# Catalog Service Management API

* [English](README.md)
* [简体中文](README_zh.md)

本项目是一个微服务 API 管理平台的后端实现，用户可以在前端 Dashboard 管理服务和版本。

项目目前包含以下功能：
- Service 的 List 和 Get 接口，支持搜索、过滤、排序、分页、查看详情等功能。
- 基于 API Key 的简单认证机制，基于角色的用户授权机制。
- 支持使用内存和 PostgreSQL 两种存储引擎。
- 包含测试代码

文件目录：



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
	curl -H "Authorization: Bearer dummy_token" http://localhost:8080/v1/services
	
	# 测试获取特定服务详情
	curl -H "Authorization: Bearer dummy_token" http://localhost:8080/v1/services/1
	```

## 开发

make 命令：
1. build production 会设置变量 `export GIN_MODE=release`
