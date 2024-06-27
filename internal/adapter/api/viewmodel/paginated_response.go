package viewmodel

// PaginatedResponse 用于分页响应的通用模型
// @Description 分页响应模型，包含分页数据和分页信息
type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	TotalCount int         `json:"total_count" example:"100"`
	Page       int         `json:"page" example:"1"`
	PageSize   int         `json:"page_size" example:"20"`
	TotalPages int         `json:"total_pages" example:"5"`
}

func NewPaginatedResponse(data interface{}, totalCount, page, pageSize int) PaginatedResponse {
	return PaginatedResponse{
		Data:       data,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: (totalCount + pageSize - 1) / pageSize,
	}
}
