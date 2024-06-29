# 设置 Go 命令
GO := go

# 设置覆盖率输出文件
COVERAGE_OUTPUT := coverage.out

# 设置 HTML 格式的覆盖率报告文件
COVERAGE_HTML := coverage.html

# 声明所有 PHONY 目标
.PHONY: run-local run-docker run-all gen-doc test test-coverage clean

# 运行本地服务
run-local:
	./scripts/run_local.sh

# 运行 Docker 服务
run-docker:
	./scripts/run_docker.sh

# 运行所有服务
run-all:
	./scripts/run_all.sh

# 生成 API 文档
gen-doc:
	swag init -g cmd/server/main.go -o api

# 运行测试
test:
	$(GO) test -v ./cmd/... ./internal/...

# 运行测试并生成覆盖率报告
test-coverage:
	$(GO) test -v -coverprofile=$(COVERAGE_OUTPUT) ./cmd/... ./internal/...
	$(GO) tool cover -html=$(COVERAGE_OUTPUT) -o $(COVERAGE_HTML)
	$(GO) tool cover -func=$(COVERAGE_OUTPUT)
	@echo "Coverage report generated: $(COVERAGE_HTML)"
	open $(COVERAGE_HTML)

# 清理生成的文件
clean:
	rm -f $(COVERAGE_OUTPUT) $(COVERAGE_HTML)
