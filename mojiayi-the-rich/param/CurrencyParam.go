package param

/**
* 纸币重量计算参数
 */
type CurrencyParam struct {
	BaseParam
	/**
	* 货币名称，比如rmb-人民币，usd-美元
	 */
	name string
	/**
	* 货币金额，比如10000元
	 */
	amount int64
	/**
	* 货币单位，比如100表示100元一张，50表示50元一张
	 */
	unit int64
}

func (p *CurrencyParam) GetName() string {
	return p.name
}

func (p *CurrencyParam) SetName(name string) *CurrencyParam {
	p.name = name
	return p
}

func (p *CurrencyParam) GetAmount() int64 {
	return p.amount
}

func (p *CurrencyParam) SetAmount(amount int64) *CurrencyParam {
	p.amount = amount
	return p
}

func (p *CurrencyParam) GetUnit() int64 {
	return p.unit
}

func (p *CurrencyParam) SetUnit(unit int64) *CurrencyParam {
	p.unit = unit
	return p
}
