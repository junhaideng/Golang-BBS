package model

type JwcNotice struct {
	Id      uint   `gorm:"AUTO_INCREMENT"`
	Title   string `gorm:"type:varchar(100)"`
	Link    string `gorm:"varchar(100)"`
	PubDate string `gorm:"varchar(30)"`
}
