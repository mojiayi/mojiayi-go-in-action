package dao

import (
	"fmt"
	"mojiayi-the-rich/setting"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID         uint64    `gorm:"primary_key" json:"id"`
	CreateTime time.Time `json: "create_time"`
	UpdateTime time.Time `json: "update_time"`
	DeleteFlag uint8     `json: "delete_flag"`
}

var DB *gorm.DB

func Setup() {
	var err error
	var dbUrl = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.MySQLSetting.User,
		setting.MySQLSetting.Password,
		setting.MySQLSetting.IP,
		setting.MySQLSetting.Port,
		setting.MySQLSetting.Database)
	DB, err = gorm.Open(mysql.Open(dbUrl), &gorm.Config{})

	if err != nil {
		fmt.Println("models setup err:", err)
	}
}
