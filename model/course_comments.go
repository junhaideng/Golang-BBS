package model

type CourseComment struct {
	Id       uint   `gorm:"AUTO_INCREMENT" json:"id"`
	Comment  string `gorm:"type:text" json:"comment"`
	CourseId uint   `gorm:"type:bigint" json:"course_id"`
	Username string `gorm:"type:varchar(40)" json:"username"`
}
