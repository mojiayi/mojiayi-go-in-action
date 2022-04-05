package mapper

import (
	"log"
	"mojiayi-the-rich/dao/domain"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewCurrencyInfoDao(DB *gorm.DB) {
	db = DB
}

func SelectByCurrencyCode(currencyCode string, nominalValue decimal.Decimal) (currencyInfo domain.CurrencyInfo, err error) {
	rows, err := db.Raw("select * from currency_info where currency_code=? and nominal_value=?", currencyCode, nominalValue).Rows()
	var record domain.CurrencyInfo
	if rows == nil {
		log.Fatal("查询失败")
		return record, nil
	}
	defer rows.Close()
	for rows.Next() {
		db.ScanRows(rows, &record)
	}

	return record, nil
}
