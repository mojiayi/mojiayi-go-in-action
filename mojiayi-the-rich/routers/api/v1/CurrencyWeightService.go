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

func CalculateWeight(context *gin.Context) {
	currencyCode := context.Query("currencyCode")
	pass, errMsg := validations.NotEmpty(currencyCode, "货币代号")
	if !pass {
		utils.IllegalArgumentErrorResp(errMsg, context)
		return
	}
	setting.MyLogger.Info("计算", currencyCode, "的重量")

	amountStr := context.Query("amount")
	pass, errMsg = validations.GreaterThanZero(amountStr, "货币金额")
	if !pass {
		utils.IllegalArgumentErrorResp(errMsg, context)
		return
	}

	nominalValueStr := context.Query("nominalValue")
	pass, errMsg = validations.GreaterThanZero(nominalValueStr, "货币单位")
	if !pass {
		utils.IllegalArgumentErrorResp(errMsg, context)
		return
	}

	amount, _ := decimal.NewFromString(amountStr)
	nominalValue, _ := decimal.NewFromString(nominalValueStr)

	var currencyParam = *new(param.CurrencyParam)
	currencyParam.SetCurrencyCode(strings.ToUpper(currencyCode))
	currencyParam.SetAmount(amount)
	currencyParam.SetNominalValue(nominalValue)
	currencyParam.SetTimestamp(int64(time.Millisecond))
	// 以下2个字段，与业务本身无关，只是为了查看访问来源才加的
	currencyParam.SetClientAgent(context.Request.UserAgent())
	currencyParam.SetClientIP(context.ClientIP())

	currencyWeightVO, err := calculateWeight(currencyParam)
	if err != nil {
		utils.ErrorResp(http.StatusGone, "计算失败，请重试！", context)
		return
	}
	utils.SuccessResp(currencyWeightVO, context)
}

func calculateWeight(param param.CurrencyParam) (currencyWeightVO vo.CurrencyWeightVO, err error) {
	currencyInfo, err := mapper.SelectByCurrencyCode(param.GetCurrencyCode(), param.GetNominalValue())
	data := new(vo.CurrencyWeightVO)
	if err != nil {
		return *data, err
	}

	pieceCount := param.GetAmount().Div(param.GetNominalValue())

	data.CurrencyCode = param.GetCurrencyCode()
	data.Amount = param.GetAmount()
	data.CurrencyName = currencyInfo.CurrencyName
	data.NominalValue = currencyInfo.NominalValue
	data.WeightInGram = currencyInfo.WeightInGram.Mul(pieceCount)
	data.WeightInKiloGram = data.WeightInGram.Div(constants.ONE_THOUSAND)
	data.WeightInTon = data.WeightInGram.Div(constants.ONE_THOUSAND).Div(constants.ONE_THOUSAND)
	data.WeightInPound = data.WeightInGram.Div(constants.ONE_THOUSAND).Mul(decimal.NewFromFloat(2.204))
	return *data, nil
}
