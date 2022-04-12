package domain

import (
	"github.com/shopspring/decimal"
)

type CurrencyInfo struct {
	BaseModel

	CurrencyCode string          `json:"currency_code"`
	CurrencyType int             `json:"currency_type"`
	CurrencyName string          `json:"currency_name"`
	NominalValue decimal.Decimal `json:"nominal_value"`
	WeightInGram decimal.Decimal `json:"weight_in_gram"`
}

type CurrencyQueryInfo struct {
	PageInfo

	CurrencyCode string `json:"currency_code"`
}
