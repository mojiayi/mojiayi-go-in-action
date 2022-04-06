package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	LogOutputDir string
)

func LoadLogOutputConfig(cfg *ini.File) {
	dir := cfg.Section("log").Key("dir").String()
	if len(dir) == 0 {
		fmt.Println("获取日志输出目录失败")
	}

	LogOutputDir = dir
}
