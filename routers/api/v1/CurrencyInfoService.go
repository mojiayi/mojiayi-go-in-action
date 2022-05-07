package v1

import (
	"mojiayi-the-rich/dao/domain"
	"mojiayi-the-rich/dao/mapper"
	"mojiayi-the-rich/utils"

	"github.com/gin-gonic/gin"
)

type CurrencyInfoService struct {
	currencyInfoMapper mapper.CurrencyInfoMapper
	paginateUtil       utils.PaginateUtil
	respUtil           utils.RespUtil
}

func (c *CurrencyInfoService) QueryAvailableCurrency(ctx *gin.Context) {
	pageResult := domain.BasePageResult{}
	pageResult.CurrentPage = c.paginateUtil.GetCurrentPage(ctx)
	pageResult.PageSize = c.paginateUtil.GetPageSize(ctx)

	currencyCode := ctx.Query("currencyCode")

	total := c.currencyInfoMapper.CountByCondition(currencyCode)
	pageResult.Total = total
	if total == 0 {
		pageResult.Pages = 0
		pageResult.Data = make(map[string]interface{}, 0)
		c.respUtil.SuccessResp(&pageResult, ctx)
		return
	}
	list, err := c.currencyInfoMapper.PageByCondition(&pageResult, currencyCode)
	if err != nil {
		pageResult.Pages = 0
		pageResult.Data = make(map[string]interface{}, 0)
		c.respUtil.SuccessResp(&pageResult, ctx)
		return
	}
	pageResult.Data = list
	c.respUtil.SuccessResp(&pageResult, ctx)
}
