package config

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/jessevdk/go-flags"
	"go.uber.org/dig"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func LoadProjectConfig() {
	container := dig.New()

	container.Provide(InitOption)
	container.Provide(InitConfig)
	container.Provide(InitPersistentConfig)
	container.Provide(InitMySQLConfig)

	err := container.Invoke(InitMySQLDatasource)
	if err != nil {
		log.Fatal(err)
	}
}

func InitOption() (*Option, error) {
	var opt Option
	_, err := flags.Parse(&opt)

	return &opt, err
}

func InitConfig(opt *Option) (*ini.File, error) {
	cfg, err := ini.Load(opt.ConfigFile)
	return cfg, err
}

func InitMySQLConfig(cfg *ini.File) (*MySQLConfig, error) {
	port, err := cfg.Section("mysql").Key("port").Int()
	if err != nil {
		return nil, err
	}

	return &MySQLConfig{
		IP:       cfg.Section("mysql").Key("ip").String(),
		Port:     port,
		User:     cfg.Section("mysql").Key("user").String(),
		Password: cfg.Section("mysql").Key("password").String(),
		Database: cfg.Section("mysql").Key("database").String(),
	}, nil
}

func InitPersistentConfig(cfg *ini.File) (*PersistentConfig, error) {
	var persistentConfig PersistentConfig

	// mysql, err := InitMySQLConfig(cfg)
	// if err != nil {
	// return &persistentConfig, err
	// }

	// persistentConfig.MySQL = mysql
	return &persistentConfig, nil
}

func InitMySQLDatasource(mysqlConfig *MySQLConfig) (*gorm.DB, error) {
	// cannot depend on result objects: *config.PersistentConfig embeds a dig.Out
	dsn := mysqlConfig.User + ":" + mysqlConfig.Password + "@tcp(" + mysqlConfig.IP + ":" + strconv.Itoa(mysqlConfig.Port) + ")/" + mysqlConfig.Database + "?charset=utf8mb4&parseTime=true&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("连接数据库失败", err)
		return DB, err
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("创建数据库连接池失败", err)
		return DB, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return DB, err
}

func PrintInfo(mysql *MySQLConfig) {
	fmt.Println("=========== mysql section ===========")
	fmt.Println("mysql ip:", mysql.IP)
	fmt.Println("mysql port:", mysql.Port)
	fmt.Println("mysql user:", mysql.User)
	fmt.Println("mysql password:", mysql.Password)
	fmt.Println("mysql db:", mysql.Database)
}
