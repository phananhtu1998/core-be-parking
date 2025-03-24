package model

// PaginationInfo chứa thông tin về phân trang
type PaginationInfo struct {
	CurrentPage int   `json:"current_page"`
	PageSize    int   `json:"page_size"`
	Total       int64 `json:"total"`
}

// PaginationResponse là struct generic cho response có phân trang
type PaginationResponse[T any] struct {
	Items      T              `json:"items"`
	Pagination PaginationInfo `json:"pagination"`
}

// NewPaginationResponse tạo một response phân trang mới
func NewPaginationResponse[T any](items T, page, pageSize int, total int64) PaginationResponse[T] {
	return PaginationResponse[T]{
		Items: items,
		Pagination: PaginationInfo{
			CurrentPage: page,
			PageSize:    pageSize,
			Total:       total,
		},
	}
}
