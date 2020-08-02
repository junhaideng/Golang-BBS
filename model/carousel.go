package model

import "time"

type Carousel struct {
	Id       uint      `gorm:"AUTO_INCREMENT"`
	Title    string    `gorm:"varchar(30)"`
	Time     time.Time `gorm:"type:datetime;default:current_timestamp"`
	Active   bool      `gorm:"type:bool;default:true"`
	Path     string    `gorm:"type:varchar(200);not null"`
	Filename string    `gorm:"type:varchar(40)"`
	Url      string    `gorm:"type:varchar(100)"`
}
