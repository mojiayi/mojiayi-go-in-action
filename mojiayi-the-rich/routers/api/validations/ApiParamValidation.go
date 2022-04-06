package validations

import (
	"strconv"
)

func NotEmpty(value string, key string) (bool, string) {
	if len(value) == 0 {
		return false, key + "不能为空"
	}
	return true, ""
}

func IsEmpty(value string, key string) (bool, string) {
	if len(value) != 0 {
		return false, key + "必须为空"
	}
	return true, ""
}

func GreaterThanZero(value string, key string) (bool, string) {
	pass, errMsg := NotEmpty(value, key)
	if !pass {
		return pass, errMsg
	}
	num, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return false, key + "必须是整数"
	}

	if num <= 0 {
		return false, key + "必须大于0"
	}
	return true, ""
}
