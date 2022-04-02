package vo

import "github.com/shopspring/decimal"

/**
* 纸币重量计算结果
 */
type CurrencyWeightVO struct {
	/**
	* 货币代号
	 */
	CurrencyCode string
	/**
	* 货币名称，比如rmb-人民币，usd-美元
	 */
	CurrencyName string
	/**
	* 货币金额，比如10000元
	 */
	Amount decimal.Decimal
	/**
	 * 货币单位，比如100表示100元一张，50表示50元一张
	 */
	NominalValue decimal.Decimal
	/**
	 * 以吨为单位的货币重量，比如1吨
	 */
	WeightInTon decimal.Decimal
	/**
	 * 以千克为单位的货币重量，比如1千克
	 */
	WeightInKiloGram decimal.Decimal
	/**
	 * 以克为单位的货币重量，比如1克
	 */
	WeightInGram decimal.Decimal
	/**
	 * 以磅为单位的货币重量，比如1磅
	 */
	WeightInPound decimal.Decimal
}
