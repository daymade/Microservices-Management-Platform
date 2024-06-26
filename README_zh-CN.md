# catalog-service-management-api

* [English](README.md)
* [简体中文](README_zh.md)

## 快速开始

```bash
go run cmd/server/main.go
```

可以使用以下 curl 命令来测试端点：

```bash
# 测试获取服务列表
curl -H "Authorization: Bearer dummy_token" http://localhost:8080/v1/services

# 测试获取特定服务详情
curl -H "Authorization: Bearer dummy_token" http://localhost:8080/v1/services/1
```

