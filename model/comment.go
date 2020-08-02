package model

// 评论
type Comment struct {
	Id       uint   `gorm:"AUTO_INCREMENT" json:"id"`
	Comment  string `gorm:"type:text" json:"comment"`
	ReplyId  uint   `gorm:"type:bigint" json:"reply_id"`
	Username string `gorm:"type:varchar(30)" json:"username"`
}
