package utils

// PaginationQuery is the standard query struct for list endpoints
type PaginationQuery struct {
	Page  int `form:"page"  binding:"omitempty,min=1"`
	Limit int `form:"limit" binding:"omitempty,min=1,max=100"`
}

func (p *PaginationQuery) Normalize() {
	if p.Page == 0 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 10
	}
}

func (p *PaginationQuery) Offset() int {
	return (p.Page - 1) * p.Limit
}

// PaginationMeta is returned alongside list responses
type PaginationMeta struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
}

// ListResponse wraps items + pagination meta
type ListResponse struct {
	Items interface{}    `json:"items"`
	Meta  PaginationMeta `json:"meta"`
}
