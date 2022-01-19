package service

import (
	"mojiayi-the-rich/param"
	"mojiayi-the-rich/vo"

	"github.com/shopspring/decimal"
)

func CalculateWeight(param param.CurrencyParam) vo.CurrencyWeightVO {
	data := new(vo.CurrencyWeightVO)
	data.SetAmount(param.GetAmount())
	data.SetName(param.GetName())
	data.SetUnit(param.GetUnit())

	onePieceWeight := decimal.NewFromFloat(1.15)
	pieceCount := decimal.NewFromInt(param.GetAmount()).Div(decimal.NewFromInt(param.GetUnit()))
	oneThousand := decimal.NewFromFloat(1000)
	weightInGram := onePieceWeight.Mul(pieceCount)
	data.SetWeightInGram(weightInGram.InexactFloat64())
	data.SetWeightInKiloGram(weightInGram.Div(oneThousand).InexactFloat64())
	data.SetWeightInTon(weightInGram.Div(oneThousand).Div(oneThousand).InexactFloat64())
	data.SetWeightInPound(weightInGram.Div(oneThousand).Mul(decimal.NewFromFloat(2.204)).InexactFloat64())
	return *data
}
