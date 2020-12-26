package model

import "time"

type User struct {
	Id          uint      `gorm:"AUTO_INCREMENT" json:"-"`
	Username    string    `gorm:"type:varchar(40);not null;unique" json:"username"`
	Password    string    `gorm:"type:varchar(100);not null" json:"-"`
	Academy     string    `gorm:"type:varchar(30)" json:"academy"`
	Avatar      string    `gorm:"type:varchar(200)" json:"avatar"`
	Email       string    `gorm:"type:varchar(30)" json:"email"`
	CreateTime  time.Time `gorm:"type:datetime;default:current_timestamp" json:"-"`
	Grade       string    `gorm:"type:varchar(30)" json:"grade"`
	Age         uint8     `gorm:"type:smallint" json:"age"`
	Gender      string    `gorm:"type:varchar(30)" json:"gender"`
	Description string    `gorm:"type:text" json:"description"`
}
