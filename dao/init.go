package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

const (
	ErrorNoPermission = "无相应权限"
)

var DB *gorm.DB

func init() {
	var err error
	username := viper.GetString("database.mysql.username")
	password := viper.GetString("database.mysql.password")
	database := viper.GetString("database.mysql.database")
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, database))
	if err != nil {
		panic(err)
	}
	DB.LogMode(true)
}
