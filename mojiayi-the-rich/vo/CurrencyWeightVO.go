package vo

/**
* 纸币重量计算结果
 */
type CurrencyWeightVO struct {
	/**
	* 货币名称，比如rmb-人民币，usd-美元
	 */
	Name string
	/**
	* 货币金额，比如10000元
	 */
	Amount int64
	/**
	 * 货币单位，比如100表示100元一张，50表示50元一张
	 */
	Unit int64
	/**
	 * 以吨为单位的货币重量，比如1吨
	 */
	WeightInTon float64
	/**
	 * 以千克为单位的货币重量，比如1千克
	 */
	WeightInKiloGram float64
	/**
	 * 以克为单位的货币重量，比如1克
	 */
	WeightInGram float64
	/**
	 * 以磅为单位的货币重量，比如1磅
	 */
	WeightInPound float64
}

func (p *CurrencyWeightVO) GetName() string {
	return p.Name
}

func (p *CurrencyWeightVO) SetName(name string) *CurrencyWeightVO {
	p.Name = name
	return p
}

func (p *CurrencyWeightVO) GetAmount() int64 {
	return p.Amount
}

func (p *CurrencyWeightVO) SetAmount(amount int64) *CurrencyWeightVO {
	p.Amount = amount
	return p
}

func (p *CurrencyWeightVO) GetUnit() int64 {
	return p.Unit
}

func (p *CurrencyWeightVO) SetUnit(unit int64) *CurrencyWeightVO {
	p.Unit = unit
	return p
}

func (p *CurrencyWeightVO) GetWeightInTon() float64 {
	return p.WeightInTon
}

func (p *CurrencyWeightVO) SetWeightInTon(WeightInTon float64) *CurrencyWeightVO {
	p.WeightInTon = WeightInTon
	return p
}

func (p *CurrencyWeightVO) GetWeightInKiloGram() float64 {
	return p.WeightInKiloGram
}

func (p *CurrencyWeightVO) SetWeightInKiloGram(WeightInKiloGram float64) *CurrencyWeightVO {
	p.WeightInKiloGram = WeightInKiloGram
	return p
}

func (p *CurrencyWeightVO) GetWeightInGram() float64 {
	return p.WeightInGram
}

func (p *CurrencyWeightVO) SetWeightInGram(WeightInGram float64) *CurrencyWeightVO {
	p.WeightInGram = WeightInGram
	return p
}

func (p *CurrencyWeightVO) GetWeightInPound() float64 {
	return p.WeightInPound
}

func (p *CurrencyWeightVO) SetWeightInPound(WeightInPound float64) *CurrencyWeightVO {
	p.WeightInPound = WeightInPound
	return p
}
