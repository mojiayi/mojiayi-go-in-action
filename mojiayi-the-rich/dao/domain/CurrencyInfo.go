package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type CurrencyInfo struct {
	ID           uint64
	CurrencyCode string
	CurrencyName string
	NominalValue decimal.Decimal
	WeightInGram decimal.Decimal
	CreateTime   time.Time
	UpdateTime   time.Time
	DeleteFlag   uint8
}
