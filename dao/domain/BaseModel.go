package domain

import "time"

type BaseModel struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	DeleteFlag uint8     `json:"delete_flag"`
}

type BasePageResult struct {
	CurrentPage int         `json:"currentPage"`
	PageSize    int         `json:"pageSize"`
	Total       int         `json:"total"`
	Pages       int         `json:"pages"`
	Data        interface{} `json:"data"`
}

type PageInfo struct {
	CurrentPage int `json:"currentPage"`
	PageSize    int `json:"pageSize"`
}
