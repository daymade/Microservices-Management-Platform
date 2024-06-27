package viewmodel

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	TotalCount int         `json:"total_count"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
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
