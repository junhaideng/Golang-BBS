package model

type Reply struct {
	Id        uint   `gorm:"AUTO_INCREMENT"`
	ArticleId uint   `gorm:"type:bigint"`
	Comments  uint   `gorm:"type:bigint;default:0"`
	Reply     string `gorm:"type:text"`
	Username  string `gorm:"type:varchar(40)"`
	Dislike   uint   `gorm:"type:bigint;default:0"`
	Like      uint   `gorm:"type:bigint;default:0"`
	star      uint   `gorm:"type:bigint;default:0"`
}
