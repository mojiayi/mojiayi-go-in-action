package v1

import (
	"github.com/gin-gonic/gin"
	"mojiayi-the-rich/dao"
	"mojiayi-the-rich/dao/mapper"
	"mojiayi-the-rich/utils"
)

type CurrencyInfoService struct {
}

var currencyInfo mapper.CurrencyInfo

func QueryAvailableCurrency(ctx *gin.Context) {
	pageResult := dao.BasePageResult{}
	pageResult.CurrentPage = utils.GetCurrentPage(ctx)
	pageResult.PageSize = utils.GetPageSize(ctx)

	currencyCode := ctx.Query("currencyCode")

	total := currencyInfo.CountByCondition(currencyCode)
	pageResult.Total = total
	if total == 0 {
		pageResult.Pages = 0
		pageResult.Data = make(map[string]interface{}, 0)
		utils.SuccessResp(&pageResult, ctx)
		return
	}
	list, err := currencyInfo.PageByCondition(&pageResult, currencyCode)
	if err != nil {
		pageResult.Pages = 0
		pageResult.Data = make(map[string]interface{}, 0)
		utils.SuccessResp(&pageResult, ctx)
		return
	}
	pageResult.Data = list
	utils.SuccessResp(&pageResult, ctx)
}
