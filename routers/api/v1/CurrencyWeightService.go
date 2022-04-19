package v1

import (
	"mojiayi-the-rich/constants"
	"mojiayi-the-rich/dao/mapper"
	"mojiayi-the-rich/param"
	"mojiayi-the-rich/routers/api/validations"
	"mojiayi-the-rich/setting"
	"mojiayi-the-rich/utils"
	"mojiayi-the-rich/vo"
	"net/http"
	"strings"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type CurrencyWeightService struct {
	apiParamValidation validations.ApiParamValidation
	respUtil           utils.RespUtil
	currencyInfoMapper mapper.CurrencyInfoMapper
}

func (c *CurrencyWeightService) CalculateWeight(context *gin.Context) {
	currencyCode := context.Query("currencyCode")
	pass, errMsg := c.apiParamValidation.NotEmpty(currencyCode, "货币代号")
	if !pass {
		c.respUtil.IllegalArgumentErrorResp(errMsg, context)
		return
	}

	amountStr := context.Query("amount")
	pass, errMsg = c.apiParamValidation.GreaterThanZero(amountStr, "货币金额")
	if !pass {
		c.respUtil.IllegalArgumentErrorResp(errMsg, context)
		return
	}

	nominalValueStr := context.Query("nominalValue")
	pass, errMsg = c.apiParamValidation.GreaterThanZero(nominalValueStr, "货币单位")
	if !pass {
		c.respUtil.IllegalArgumentErrorResp(errMsg, context)
		return
	}

	amount, _ := decimal.NewFromString(amountStr)
	nominalValue, _ := decimal.NewFromString(nominalValueStr)

	setting.MyLogger.Info("计算货币重量,currencyCode=", currencyCode, ",nominalValue=", nominalValue)

	var currencyParam = *new(param.CurrencyParam)
	currencyParam.SetCurrencyCode(strings.ToUpper(currencyCode))
	currencyParam.SetAmount(amount)
	currencyParam.SetNominalValue(nominalValue)
	currencyParam.SetTimestamp(int64(time.Millisecond))
	// 以下2个字段，与业务本身无关，只是为了查看访问来源才加的
	currencyParam.SetClientAgent(context.Request.UserAgent())
	currencyParam.SetClientIP(context.ClientIP())

	currencyWeightVO, err := c.calculateWeight(currencyParam)
	if err != nil {
		c.respUtil.ErrorResp(http.StatusGone, err.Error(), context)
		return
	}
	c.respUtil.SuccessResp(currencyWeightVO, context)
}

func (c *CurrencyWeightService) calculateWeight(param param.CurrencyParam) (currencyWeightVO vo.CurrencyWeightVO, err error) {
	record, err := c.currencyInfoMapper.SelectByCurrencyCode(param.GetCurrencyCode(), param.GetNominalValue())
	data := new(vo.CurrencyWeightVO)
	if err != nil {
		setting.MyLogger.Info("货币不存在,currencyCode=", param.GetCurrencyCode(), ",nominalValue=", param.GetNominalValue())
		return *data, err
	}

	pieceCount := param.GetAmount().Div(param.GetNominalValue())

	data.CurrencyCode = param.GetCurrencyCode()
	data.Amount = param.GetAmount()
	data.CurrencyName = record.CurrencyName
	data.NominalValue = record.NominalValue
	data.WeightInGram = record.WeightInGram.Mul(pieceCount)
	data.WeightInKiloGram = data.WeightInGram.Div(constants.OneThousand)
	data.WeightInTon = data.WeightInGram.Div(constants.OneThousand).Div(constants.OneThousand)
	data.WeightInPound = data.WeightInGram.Div(constants.OneThousand).Mul(decimal.NewFromFloat(2.204))
	return *data, nil
}
