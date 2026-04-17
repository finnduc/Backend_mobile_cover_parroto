package database

import (
	"gorm.io/gorm"
)

type Query struct {
	Filters map[string]any
	Page    int
	Limit   int
	OrderBy string
}

func NewQuery() *Query {
	return &Query{
		Filters: make(map[string]any),
		Page:    1,
		Limit:   10,
	}
}

func (q *Query) SetFilter(key string, value any) *Query {
	q.Filters[key] = value
	return q
}

func (q *Query) SetPage(page int) *Query {
	if page > 0 {
		q.Page = page
	}
	return q
}

func (q *Query) SetLimit(limit int) *Query {
	if limit > 0 {
		q.Limit = limit
	}
	return q
}

func (q *Query) SetOrderBy(field string) *Query {
	q.OrderBy = field
	return q
}

func (q *Query) Apply(db *gorm.DB) *gorm.DB {
	result := db

	for key, value := range q.Filters {
		if value != nil {
			result = result.Where(key+" = ?", value)
		}
	}

	if q.OrderBy != "" {
		result = result.Order(q.OrderBy)
	}

	offset := (q.Page - 1) * q.Limit
	result = result.Offset(offset).Limit(q.Limit)

	return result
}

func (q *Query) Count(db *gorm.DB) *gorm.DB {
	result := db

	for key, value := range q.Filters {
		if value != nil {
			result = result.Where(key+" = ?", value)
		}
	}

	return result
}
