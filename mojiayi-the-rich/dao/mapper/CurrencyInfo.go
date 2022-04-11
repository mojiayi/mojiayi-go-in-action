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

func SelectByCurrencyCode(wrapper map[string]interface{}) (currencyInfo CurrencyInfo, err error) {
	var record CurrencyInfo
	dao.DB.Model(&CurrencyInfo{}).Where(wrapper).Find(&record)

	return record, nil
}

func (c *CurrencyInfo) CountByCondition(wrapper map[string]interface{}) int32 {
	var total int64
	dao.DB.Model(&CurrencyInfo{}).Where(wrapper).Count(&total)

	return int32(total)
}

func (c *CurrencyInfo) PageByCondition(pageResult *dao.BasePageResult, wrapper map[string]interface{}) (list interface{}, err error) {
	list = []CurrencyInfo{}
	err = dao.DB.Model(&CurrencyInfo{}).Scopes(dao.Paginate(pageResult)).Where(wrapper).Find(&list).Error
	return list, err
}
