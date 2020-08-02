package model

import "time"

type LoginLog struct {
	Id         uint      `gorm:"AUTO_INCREMENT"`
	Address    string    `gorm:"type:varchar(50)"`
	CreateTime time.Time `gorm:"type:datetime;default:current_timestamp"`
	Ip         string    `gorm:"type:varchar(20);not null"`
	Username   string    `gorm:"type:varchar(40);not null"`
}
