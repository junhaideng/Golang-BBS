package model

type JwcNotice struct {
	Id      uint   `gorm:"AUTO_INCREMENT" json:"id"`
	Title   string `gorm:"type:varchar(100)" json:"title"`
	Link    string `gorm:"varchar(100)" json:"link"`
	PubDate string `gorm:"varchar(30)" json:"pub_date"`
}
