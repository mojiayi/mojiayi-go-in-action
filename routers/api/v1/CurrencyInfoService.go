package v1

import (
	"github.com/gin-gonic/gin"
	"mojiayi-the-rich/dao/domain"
	"mojiayi-the-rich/dao/mapper"
	"mojiayi-the-rich/utils"
)

type CurrencyInfoService struct {
}

var currencyInfoMapper mapper.CurrencyInfoMapper
var paginateUtil utils.PaginateUtil

func (c *CurrencyInfoService) QueryAvailableCurrency(ctx *gin.Context) {
	pageResult := domain.BasePageResult{}
	pageResult.CurrentPage = paginateUtil.GetCurrentPage(ctx)
	pageResult.PageSize = paginateUtil.GetPageSize(ctx)

	currencyCode := ctx.Query("currencyCode")

	total := currencyInfoMapper.CountByCondition(currencyCode)
	pageResult.Total = total
	if total == 0 {
		pageResult.Pages = 0
		pageResult.Data = make(map[string]interface{}, 0)
		utils.SuccessResp(&pageResult, ctx)
		return
	}
	list, err := currencyInfoMapper.PageByCondition(&pageResult, currencyCode)
	if err != nil {
		pageResult.Pages = 0
		pageResult.Data = make(map[string]interface{}, 0)
		utils.SuccessResp(&pageResult, ctx)
		return
	}
	pageResult.Data = list
	utils.SuccessResp(&pageResult, ctx)
}
