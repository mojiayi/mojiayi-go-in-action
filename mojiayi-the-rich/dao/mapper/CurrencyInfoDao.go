package mapper

import (
	"log"
	"mojiayi-the-rich/dao/domain"

	"github.com/shopspring/decimal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SelectByCurrencyCode(currencyCode string, nominalValue decimal.Decimal) (currencyInfo domain.CurrencyInfo, err error) {
	dsn := "root:123456@tcp(localhost:3306)/sys?charset=utf8mb4&parseTime=true&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	var record domain.CurrencyInfo
	if err != nil {
		log.Fatal("连接数据库失败", err)
		return record, err
	}

	DB.Table("currency_info").Where("currency_code=? and nominal_value=?", currencyCode, nominalValue).First(&record)
	return record, nil
}
