package model

type Feedback struct {
	Id       uint   `gorm:"AUTO_INCREMENT"`
	Content  string `gorm:"type:text;not null"`
	Email    string `gorm:"type:varchar(30)"`
	Title    string `gorm:"type:varchar(200)"`
	Username string `gorm:"type:varchar(40)"`
	Active   bool   `gorm:"type:bool;default:true"`
}
