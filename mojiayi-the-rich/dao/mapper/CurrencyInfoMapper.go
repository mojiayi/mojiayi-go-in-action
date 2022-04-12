package mapper

import (
	"errors"
	"github.com/shopspring/decimal"
	"mojiayi-the-rich/dao/domain"
	"mojiayi-the-rich/setting"
	"mojiayi-the-rich/utils"
)

type CurrencyInfoMapper struct {
}

var paginateUtil utils.PaginateUtil

func (c *CurrencyInfoMapper) SelectByCurrencyCode(currencyCode string, nominal decimal.Decimal) (currencyInfo domain.CurrencyInfo, err error) {
	wrapper := make(map[string]interface{})
	wrapper["currency_code"] = currencyCode
	wrapper["nominal_value"] = nominal

	var record domain.CurrencyInfo
	setting.DB.Model(&domain.CurrencyInfo{}).Where(wrapper).Find(&record)

	if record.ID == 0 {
		return record, errors.New("货币" + currencyCode + "(" + nominal.String() + ")不存在")
	}

	return record, nil
}

func (c *CurrencyInfoMapper) CountByCondition(currencyCode string) int {
	wrapper := make(map[string]interface{}, 0)
	if currencyCode != "" {
		wrapper["currency_code"] = currencyCode
	}

	var total int64
	setting.DB.Model(&domain.CurrencyInfo{}).Where(wrapper).Count(&total)

	return int(total)
}

func (c *CurrencyInfoMapper) PageByCondition(pageResult *domain.BasePageResult, currencyCode string) (list []domain.CurrencyInfo, err error) {
	wrapper := make(map[string]interface{}, 0)
	if currencyCode != "" {
		wrapper["currency_code"] = currencyCode
	}

	list = []domain.CurrencyInfo{}
	err = setting.DB.Model(&domain.CurrencyInfo{}).Scopes(paginateUtil.Paginate(pageResult)).Where(wrapper).Find(&list).Error
	return list, err
}
