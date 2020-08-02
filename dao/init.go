package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "Edgar:Edgar@/bbs?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DB.LogMode(true)
}
