package utils

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mojiayi-the-rich/dao/domain"
	"strconv"
)

type PaginateUtil struct {
}

var (
	DefaultCurrentPage int = 1
	DefaultPageSize    int = 10
	MaxPageSize        int = 200
)

func (p *PaginateUtil) GetCurrentPage(ctx *gin.Context) int {
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

func (p *PaginateUtil) GetPageSize(ctx *gin.Context) int {
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

func (p *PaginateUtil) Paginate(page *domain.BasePageResult) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page.Pages = page.Total / page.PageSize
		if page.Total%page.PageSize != 0 {
			page.Pages++
		}
		if page.CurrentPage > page.Pages {
			page.CurrentPage = page.Pages
		}
		offset := int((page.CurrentPage - 1) * page.PageSize)
		return db.Offset(offset).Limit(int(page.PageSize))
	}
}
