package model

type Course struct {
	Id         uint   `gorm:"AUTO_INCREMENT" json:"id"`
	School     string `gorm:"type:varchar(60)" json:"school"`
	CourseName string `gorm:"type:varchar(40)" json:"course_name"`
	Teacher    string `gorm:"type:varchar(30)" json:"teacher"`
	Type       string `gorm:"type:varchar(30)" json:"type"`
}
