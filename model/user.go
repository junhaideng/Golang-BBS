package model

import "time"

type User struct {
	Id          uint      `gorm:"AUTO_INCREMENT"`
	Username    string    `gorm:"type:varchar(40);not null;unique"`
	Password    string    `gorm:"type:varchar(100);not null"`
	Academy     string    `gorm:"type:varchar(30)"`
	Avatar      string    `gorm:"type:varchar(200)"`
	Email       string    `gorm:"type:varchar(30)"`
	CreateTime  time.Time `gorm:"type:datetime;default:current_timestamp"`
	Grade       string    `gorm:"type:varchar(30)"`
	Age         uint8     `gorm:"type:smallint"`
	Gender      string    `gorm:"type:varchar(30)"`
	Description string    `gorm:"type:text"`
}
