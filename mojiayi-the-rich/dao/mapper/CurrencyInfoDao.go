package mapper

import (
	"log"
	"mojiayi-the-rich/config"
	"mojiayi-the-rich/dao/domain"

	"github.com/shopspring/decimal"
)

func SelectByCurrencyCode(currencyCode string, nominalValue decimal.Decimal) (currencyInfo domain.CurrencyInfo, err error) {
	rows, err := config.DB.Raw("select * from currency_info where currency_code=? and nominal_value=?", currencyCode, nominalValue).Rows()
	var record domain.CurrencyInfo
	if rows == nil {
		log.Fatal("查询失败")
		return record, nil
	}
	defer rows.Close()
	for rows.Next() {
		config.DB.ScanRows(rows, &record)
	}

	return record, nil
}
