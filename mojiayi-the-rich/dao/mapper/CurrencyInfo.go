package mapper

import (
	"mojiayi-the-rich/dao"
	"mojiayi-the-rich/enums"

	"github.com/shopspring/decimal"
)

type CurrencyInfo struct {
	dao.BaseModel

	CurrencyCode string          `json: "currency_code"`
	CurrencyName string          `json: "currency_name"`
	NominalValue decimal.Decimal `json: "nominal_value"`
	WeightInGram decimal.Decimal `json: "weight_in_gram"`
}

func SelectByCurrencyCode(currencyCode string, nominalValue decimal.Decimal) (currencyInfo CurrencyInfo, err error) {
	rows, _ := dao.DB.Raw("select * from currency_info where currency_code=? and nominal_value=? and delete_flag=?", currencyCode, nominalValue, enums.NORMAL).Rows()
	var record CurrencyInfo

	defer rows.Close()
	for rows.Next() {
		dao.DB.ScanRows(rows, &record)
	}

	return record, nil
}
