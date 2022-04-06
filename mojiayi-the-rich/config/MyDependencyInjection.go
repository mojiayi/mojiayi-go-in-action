package config

import "mojiayi-the-rich/dao/mapper"

func SetupDependencyInjection() {
	container := LoadProjectConfig()

	container.Invoke(mapper.NewCurrencyInfoDao)

	container.Invoke(LoadLogOutputConfig)
}
