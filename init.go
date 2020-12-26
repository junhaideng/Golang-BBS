package main

import (
	"bbs/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

// 初始化数据库表
func InitDB() {
	username := viper.GetString("database.mysql.username")
	password := viper.GetString("database.mysql.password")
	database := viper.GetString("database.mysql.database")

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, database))
	defer db.Close()
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.AutoMigrate(&model.Avatar{})
	db.AutoMigrate(&model.Article{})
	db.AutoMigrate(&model.Carousel{})
	db.AutoMigrate(&model.Comment{})
	db.AutoMigrate(&model.Course{})
	db.AutoMigrate(&model.CourseComment{})
	db.AutoMigrate(&model.Feedback{})
	db.AutoMigrate(&model.File{})
	db.AutoMigrate(&model.JwcNotice{})
	db.AutoMigrate(&model.LoginLog{})
	db.AutoMigrate(&model.Message{})
	db.AutoMigrate(&model.MessageSetting{})
	db.AutoMigrate(&model.Reply{})
	db.AutoMigrate(&model.User{})
}
