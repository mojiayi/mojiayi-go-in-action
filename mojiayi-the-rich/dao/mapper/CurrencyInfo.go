package mapper

import (
	"github.com/shopspring/decimal"
	"mojiayi-the-rich/dao"
)

type CurrencyInfo struct {
	dao.BaseModel

	CurrencyCode string          `json:"currency_code"`
	CurrencyType int32           `json:"currency_type"`
	CurrencyName string          `json:"currency_name"`
	NominalValue decimal.Decimal `json:"nominal_value"`
	WeightInGram decimal.Decimal `json:"weight_in_gram"`
}

type CurrencyQueryInfo struct {
	dao.PageInfo

	CurrencyCode string `json:"currency_code"`
}

func SelectByCurrencyCode(currencyCode string, nominal decimal.Decimal) (currencyInfo CurrencyInfo, err error) {
	wrapper := make(map[string]interface{})
	wrapper["currency_code"] = currencyCode
	wrapper["nominal_value"] = nominal

	var record CurrencyInfo
	dao.DB.Model(&CurrencyInfo{}).Where(wrapper).Find(&record)

	return record, nil
}

func (c *CurrencyInfo) CountByCondition(currencyCode string) int32 {
	wrapper := make(map[string]interface{}, 0)
	if currencyCode != "" {
		wrapper["currency_code"] = currencyCode
	}

	var total int64
	dao.DB.Model(&CurrencyInfo{}).Where(wrapper).Count(&total)

	return int32(total)
}

func (c *CurrencyInfo) PageByCondition(pageResult *dao.BasePageResult, currencyCode string) (list interface{}, err error) {
	wrapper := make(map[string]interface{}, 0)
	if currencyCode != "" {
		wrapper["currency_code"] = currencyCode
	}

	list = []CurrencyInfo{}
	err = dao.DB.Model(&CurrencyInfo{}).Scopes(dao.Paginate(pageResult)).Where(wrapper).Find(&list).Error
	return list, err
}
