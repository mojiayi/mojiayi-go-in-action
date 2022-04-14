package param

import "github.com/shopspring/decimal"

/**
* 纸币重量计算参数
 */
type CurrencyParam struct {
	BaseParam
	/**
	* 货币名称，比如rmb-人民币，usd-美元
	 */
	currencyCode string
	/**
	* 货币金额，比如10000元
	 */
	amount decimal.Decimal
	/**
	* 货币面值，比如100表示100元一张，50表示50元一张
	 */
	nominalValue decimal.Decimal
}

func (p *CurrencyParam) GetCurrencyCode() string {
	return p.currencyCode
}

func (p *CurrencyParam) SetCurrencyCode(name string) *CurrencyParam {
	p.currencyCode = name
	return p
}

func (p *CurrencyParam) GetAmount() decimal.Decimal {
	return p.amount
}

func (p *CurrencyParam) SetAmount(amount decimal.Decimal) *CurrencyParam {
	p.amount = amount
	return p
}

func (p *CurrencyParam) GetNominalValue() decimal.Decimal {
	return p.nominalValue
}

func (p *CurrencyParam) SetNominalValue(unit decimal.Decimal) *CurrencyParam {
	p.nominalValue = unit
	return p
}
