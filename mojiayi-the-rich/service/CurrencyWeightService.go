package service

import (
	"mojiayi-the-rich/constants"
	"mojiayi-the-rich/dao/mapper"
	"mojiayi-the-rich/param"
	"mojiayi-the-rich/utils"
	"mojiayi-the-rich/vo"
	"net/http"
	"strings"

	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"github.com/go-eden/routine"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

var localTraceId = routine.NewLocalStorage()
var localContext = routine.NewLocalStorage()

func CalculateWeight(context *gin.Context) {
	localContext.Set(context)

	traceId := uuid.New()
	localTraceId.Set(traceId)

	currencyCode := context.Query("currencyCode")

	if len(currencyCode) == 0 {
		utils.IllegalArgumentErrorResp("货币代号不能为空", context)
		return
	}
	logrus.Info("计算", currencyCode, "的重量")

	amountStr := context.Query("amount")
	if len(amountStr) == 0 {
		utils.IllegalArgumentErrorResp("货币金额不能为空", context)
		return
	}

	amount, err := strconv.ParseInt(amountStr, 10, 64)
	if err != nil {
		utils.IllegalArgumentErrorResp("货币金额只能是数字", context)
		return
	}
	if amount <= 0 {
		utils.IllegalArgumentErrorResp("货币金额必须大于0", context)
		return
	}

	nominalValueStr := context.Query("nominalValue")
	if len(nominalValueStr) == 0 {
		utils.IllegalArgumentErrorResp("货币单位不能为空", context)
		return
	}

	nominalValue, err := strconv.ParseInt(nominalValueStr, 10, 64)
	if err != nil {
		utils.IllegalArgumentErrorResp("货币单位只能是数字", context)
		return
	}
	if nominalValue <= 0 {
		utils.IllegalArgumentErrorResp("货币单位必须大于0", context)
		return
	}

	var param param.CurrencyParam = *new(param.CurrencyParam)
	param.SetCurrencyCode(strings.ToUpper(currencyCode))
	param.SetAmount(decimal.NewFromInt(amount))
	param.SetNominalValue(decimal.NewFromInt(nominalValue))
	param.SetTimestamp(int64(time.Millisecond))
	// 以下2个字段，与业务本身无关，只是为了查看访问来源才加的
	param.SetClientAgent(context.Request.UserAgent())
	param.SetClientIP(context.ClientIP())

	currencyWeightVO, err := calculateWeight(param)
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
