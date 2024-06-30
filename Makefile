# 声明所有 PHONY 目标
.PHONY: run-local run-docker run-all gen-doc test test-coverage test-integration test-clean

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

# 运行单元测试
test:
	./scripts/test/run_unit_tests.sh

# 运行测试并生成覆盖率报告
test-coverage:
	./scripts/test/run_coverage_tests.sh

# 运行集成测试
test-integration:
	./scripts/test/integration/run_integration_tests.sh

# 清理测试生成的文件
test-clean:
	./scripts/test/clean_coverage_outputs.sh
