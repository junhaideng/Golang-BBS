package dao

import "bbs/model"

func GetMessageSettingsByUsername(username string) (*model.MessageSetting, error) {
	var settings = &model.MessageSetting{}
	if err := DB.Model(&model.MessageSetting{}).Where("username = ?", username).Scan(settings).Error; err != nil {
		return nil, err
	}
	return settings, nil
}

// 使用map来进行更新
// 使用struct更新的时候，会自动忽略零值
func UpdateMessageSettingsByUsername(setting map[string]interface{}) error {
	tx := DB.Begin()
	if err := tx.Model(&model.MessageSetting{}).Where("username = ?", setting["username"]).Updates(setting).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
