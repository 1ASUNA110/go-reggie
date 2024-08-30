package response

// Page 分页结构体
type Page[T any] struct {
	Total    int64 `json:"total"`    // 总记录数
	Records  []T   `json:"records"`  // 当前页的记录
	Page     int   `json:"page"`     // 当前页码
	PageSize int   `json:"pageSize"` // 每页显示的记录数
}

// NewPage 创建一个新的分页对象
func NewPage[T any](total int64, records []T, page int, pageSize int) Page[T] {
	return Page[T]{
		Total:    total,
		Records:  records,
		Page:     page,
		PageSize: pageSize,
	}
}
