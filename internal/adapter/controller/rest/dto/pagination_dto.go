package dto

type Paginated[t any] struct {
	Page       int   `json:"page"`
	PerPage    int   `json:"per_page"`
	TotalPage  int   `json:"total_page"`
	TotalCount int64 `json:"total_count"`
	Data       t     `json:"data"`
}
