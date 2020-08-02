package dao

import "bbs/model"

func FindAllFiles() ([]*model.File, error) {
	var files []*model.File
	if err := DB.Model(&model.File{}).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func FindFileByUsername(username string) ([]*model.File, error) {
	var files []*model.File
	if err := DB.Model(&model.File{}).Find(&files, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func CreateFile(file *model.File) error {
	tx := DB.Begin()
	if err := DB.Model(&model.File{}).Create(&file).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
