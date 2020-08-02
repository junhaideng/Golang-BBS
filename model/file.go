package model

import "time"

type File struct {
	Id          uint      `gorm:"AUTO_INCREMENT" json:"id"`
	Username    string    `gorm:"type:varchar(40);not null" json:"username"`
	Description string    `gorm:"type:text" json:"description"`
	Filename    string    `gorm:"type:varchar(100);not null" json:"filename"`
	Path        string    `gorm:"type:varchar(200);not null" json:"path"`
	Type        string    `gorm:"type:varchar(30);not null" json:"type"`
	UploadTime  time.Time `gorm:"type:datetime;default:current_timestamp" json:"upload_time"`
}
