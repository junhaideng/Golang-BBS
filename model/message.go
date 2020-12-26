package model

import "time"

type Message struct {
	Id       uint      `gorm:"AUTO_INCREMENT" json:"id"`
	Content  string    `gorm:"type:text" json:"content"`
	Read     bool      `gorm:"type:bool;default:false" json:"read"`
	Time     time.Time `gorm:"type:datetime;default:current_timestamp" json:"time"`
	Title    string    `gorm:"type:varchar(100)" json:"title"`
	Username string    `gorm:"type:varchar(40)" json:"-"`
	Type     string    `gorm:"type:varchar(20)" json:"type"`
	Url      string    `gorm:"type:varchar(200)" json:"url"`
}
