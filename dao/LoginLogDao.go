package dao

import "bbs/model"

func FindLoginLogByUsername(username string) ([]model.LoginLog, error) {
	var log []model.LoginLog
	if err := DB.Model(&model.LoginLog{}).Where("username = ?", username).Find(&log).Error; err != nil {
		return nil, err
	}
	return log, nil
}

func CreateLoginLog(log *model.LoginLog) error {
	tx := DB.Begin()
	if err := tx.Model(&model.LoginLog{}).Create(log).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
