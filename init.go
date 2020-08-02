package main

import (
	"bbs/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 初始化数据库表
func InitDB() {
	db, err := gorm.Open("mysql", "Edgar:Edgar@/bbs?charset=utf8&parseTime=True&loc=Local")
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
