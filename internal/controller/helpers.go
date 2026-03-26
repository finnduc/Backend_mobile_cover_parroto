package controller

import (
	"go-familytree/pkg/utils"

	"github.com/gin-gonic/gin"
)

// pagination extracts and normalizes PaginationQuery from gin context
func pagination(c *gin.Context) utils.PaginationQuery {
	var q utils.PaginationQuery
	_ = c.ShouldBindQuery(&q)
	q.Normalize()
	return q
}
