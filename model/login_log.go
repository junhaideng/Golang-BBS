package model

import "time"

type LoginLog struct {
	Id         uint      `gorm:"AUTO_INCREMENT" json:"id"`
	Address    string    `gorm:"type:varchar(50)" json:"address"`
	CreateTime time.Time `gorm:"type:datetime;default:current_timestamp" json:"create_time"`
	Ip         string    `gorm:"type:varchar(20);not null" json:"ip"`
	Username   string    `gorm:"type:varchar(40);not null" json:"username"`
}
