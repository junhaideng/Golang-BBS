package model

type MessageSetting struct {
	Id       uint   `gorm:"AUTO_INCREMENT" json:"-"`
	Username string `gorm:"type:varchar(20)" json:"-"`
	Comment  bool   `gorm:"type:bool;default:true" json:"comment"`
	Star     bool   `gorm:"type:bool;default:true" json:"star"`
	Like     bool   `gorm:"type:bool;default:true" json:"like"`
}
