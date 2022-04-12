package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	DefaultCurrentPage int = 1
	DefaultPageSize    int = 10
	MaxPageSize        int = 200
)

func GetCurrentPage(ctx *gin.Context) int {
	tempStr := ctx.Query("currentPage")
	if tempStr == "" {
		return DefaultCurrentPage
	}
	currentPage, err := strconv.Atoi(tempStr)
	if err != nil {
		return DefaultCurrentPage
	}
	if currentPage <= 0 {
		return DefaultCurrentPage
	}
	return int(currentPage)
}

func GetPageSize(ctx *gin.Context) int {
	tempStr := ctx.Query("pageSize")
	if tempStr == "" {
		return DefaultPageSize
	}
	pageSize, err := strconv.Atoi(tempStr)
	if err != nil {
		return DefaultPageSize
	}
	if pageSize <= 0 || int(pageSize) > MaxPageSize {
		return DefaultPageSize
	}
	return int(pageSize)
}
