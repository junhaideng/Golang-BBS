package model

type Reply struct {
	Id        uint   `gorm:"AUTO_INCREMENT" json:"id"`
	ArticleId uint   `gorm:"type:bigint" json:"article_id"`
	Comments  uint   `gorm:"type:bigint;default:0" json:"comments"`
	Reply     string `gorm:"type:text" json:"reply"`
	Username  string `gorm:"type:varchar(40)" json:"username"`
	Dislike   uint   `gorm:"type:bigint;default:0" json:"dislike"`
	Like      uint   `gorm:"type:bigint;default:0" json:"like"`
}
