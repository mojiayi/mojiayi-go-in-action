package dao

import (
	"fmt"
	"gorm.io/gorm/schema"
	"mojiayi-the-rich/setting"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID         uint64    `gorm:"primary_key" json:"id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	DeleteFlag uint8     `json:"delete_flag"`
}

type BasePageResult struct {
	CurrentPage int32       `json:"currentPage"`
	PageSize    int32       `json:"pageSize"`
	Total       int32       `json:"total"`
	Pages       int32       `json:"pages"`
	Data        interface{} `json:"data"`
}

type PageInfo struct {
	CurrentPage int32 `json:"currentPage"`
	PageSize    int32 `json:"pageSize"`
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
	DB, err = gorm.Open(mysql.Open(dbUrl), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})

	if err != nil {
		fmt.Println("models setup err:", err)
	}
}

func Paginate(page *BasePageResult) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page.Pages = page.Total / page.PageSize
		if page.Total%page.PageSize != 0 {
			page.Pages++
		}
		if page.CurrentPage > page.Pages {
			page.CurrentPage = page.Pages
		}
		offset := int((page.CurrentPage - 1) * page.PageSize)
		return db.Offset(offset).Limit(int(page.PageSize))
	}
}
