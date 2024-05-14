package dto

import "math"

type Paginated[t any] struct {
	Page       int   `json:"page"`
	PerPage    int   `json:"per_page"`
	TotalPage  int   `json:"total_page"`
	TotalCount int64 `json:"total_count"`
	Data       t     `json:"data"`
}

func NewPaginatedRes[t any](data t, page, perPage int, totalCount int64) Paginated[t] {
	return Paginated[t]{
		Page:       page,
		PerPage:    perPage,
		TotalPage:  int(math.Ceil(float64(totalCount) / float64(perPage))),
		TotalCount: totalCount,
		Data:       data,
	}
}
