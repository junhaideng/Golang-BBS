package model

import "time"

// 用户的帖子
type Article struct {
	Id uint `gorm:"AUTO_INCREMENT" json:"id"`
	// 哪一个用户发的
	Username string `gorm:"type:varchar(40)" json:"username"`
	// 发的标题
	Title string `gorm:"type:varchar(200)" json:"title"`
	// 发的内容
	Content string `gorm:"type:text" json:"content"`
	// 发的时间
	CreateTime time.Time `gorm:"type:datetime;default:current_timestamp" json:"create_time"`
	Type       string    `gorm:"type:varchar(40)" json:"type"`
	Star       uint      `gorm:"type:bigint;default:0" json:"star"`
	Comments   uint      `gorm:"type:bigint;default:0" json:"comments"`
	Read       uint      `gorm:"type:bigint;default:0" json:"read"`
}
