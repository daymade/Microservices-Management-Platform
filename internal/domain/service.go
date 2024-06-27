package domain

import (
	"catalog-service-management-api/internal/domain/models"
)

// Service 定义了服务管理的核心业务逻辑接口
// 这个接口提供了查询和获取服务信息的方法
type Service interface {
	// ListServices 返回服务列表，支持分页、排序和查询
	//
	// 参数:
	//   - query: 查询字符串，用于过滤服务（可以是服务名称或描述的一部分）
	//   - sortBy: 排序字段（例如："name", "created_at"）
	//   - sortDir: 排序方向（"asc" 为升序，"desc" 为降序）
	//   - page: 当前页码（从1开始）
	//   - pageSize: 每页的服务数量
	//
	// 返回值:
	//   - []models.Service: 符合条件的服务列表
	//   - int: 符合条件的服务总数
	//   - error: 如果发生错误，返回相应的错误信息
	ListServices(query string, sortBy string, sortDir string, page int, pageSize int) ([]models.Service, int, error)

	// GetService 根据ID获取特定服务的详细信息
	//
	// 参数:
	//   - id: 服务的唯一标识符
	//
	// 返回值:
	//   - models.Service: 请求的服务详细信息
	//   - error: 如果服务不存在或发生其他错误，返回相应的错误信息
	GetService(id string) (models.Service, error)
}
