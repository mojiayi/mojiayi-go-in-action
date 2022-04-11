package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	DefaultCurrentPage int32 = 1
	DefaultPageSize    int32 = 10
	MaxPageSize        int32 = 200
)

func GetCurrentPage(ctx *gin.Context) int32 {
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
	return int32(currentPage)
}

func GetPageSize(ctx *gin.Context) int32 {
	tempStr := ctx.Query("pageSize")
	if tempStr == "" {
		return DefaultPageSize
	}
	pageSize, err := strconv.Atoi(tempStr)
	if err != nil {
		return DefaultPageSize
	}
	if pageSize <= 0 || int32(pageSize) > MaxPageSize {
		return DefaultPageSize
	}
	return int32(pageSize)
}
