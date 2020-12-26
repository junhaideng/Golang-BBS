package dao

import "bbs/model"

func GetUnreadMessageByUsername(username string) ([]*model.Message, error) {
	var messages []*model.Message
	if err := DB.Model(&model.Message{}).Where("username = ? AND `read` = ?", username, false).Scan(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func GetUnreadMessageNumByUsername(username string) int {
	messages, err := GetUnreadMessageByUsername(username)
	if err != nil {
		return 0
	}
	return len(messages)
}

func GetMessageByUsername(username string) ([]*model.Message, error) {
	var messages []*model.Message
	if err := DB.Model(&model.Message{}).Where("username = ? ", username).Scan(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func SetMessageRead(username string, id string) error {
	tx := DB.Begin()
	if err := tx.Model(&model.Message{}).Where("username = ? AND id = ?", username, id).Update("read", true).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func DeleteMessage(username string, id string) error {
	tx := DB.Begin()
	if err := tx.Where("username = ? AND id = ?", username, id).Delete(&model.Message{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
