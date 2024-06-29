## `/cmd`

直接运行 `main.go` 文件，可启动后端服务（使用内存数据库，没有其他依赖）

```shell
go run server/main.go
```

启动后可访问 `http://localhost:8080` 查看效果。

```shell
curl -H "Authorization: Bearer dummy_token" http://localhost:8080/api/v1/services/1
```
