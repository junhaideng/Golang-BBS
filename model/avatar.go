package model

type Avatar struct {
	Id       uint   `gorm:"AUTO_INCREMENT"`
	Path     string `gorm:"type:varchar(200);default:'default.jpg'"`
	Username string `gorm:"type:varchar(40);not null"`
}
