## 测试策略

1. 单元测试：
	- 对核心业务逻辑进行了单元测试
	- 使用模拟对象隔离外部依赖

2. 集成测试：
	- 实现了端到端的 API 测试，确保各组件能够正确协同工作
	- 使用 Docker 环境进行测试，确保测试环境与生产环境一致

3. 测试覆盖率：
	- 使用 Go 内置的测试覆盖率工具生成报告
	- 设置了最低测试覆盖率标准，确保代码质量
