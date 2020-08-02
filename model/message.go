package model

import "time"

type Message struct {
	Id       uint      `gorm:"AUTO_INCREMENT"`
	Content  string    `gorm:"type:text"`
	Read     bool      `gorm:"type:bool;default:false"`
	Time     time.Time `gorm:"type:datetime;default:current_timestamp"`
	Title    string    `gorm:"type:varchar(100)"`
	Username string    `gorm:"type:varchar(40)"`
	Type     string    `gorm:"type:varchar(20)"`
	Url      string    `gorm:"type:varchar(200)"`
}
