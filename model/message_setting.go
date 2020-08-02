package model

type MessageSetting struct {
	Id       uint   `gorm:"AUTO_INCREMENT"`
	Username string `gorm:"type:varchar(20)"`
	Comment  bool   `gorm:"type:bool;default:true"`
	Star     bool   `gorm:"type:bool;default:true"`
	Like     bool   `gorm:"type:bool;default:true"`
}
