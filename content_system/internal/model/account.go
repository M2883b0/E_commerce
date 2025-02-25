package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Phone_number string `gorm:"column:phone_number;"`
	Password     string `gorm:"column:password;"`
	User_name    string `gorm:"column:user_name;"`
	User_type    int32  `gorm:"column:user_type;"`
	Img_url      string `gorm:"column:img_url;"`
	Description  string `gorm:"column:description;"`
}

// 指定表名
func (Account) TableName() string {
	table := "ec.user"
	return table
}
