package dao

import (
	"bbs/model"
	"errors"
	"fmt"
	"os"
)

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

func FindFileByID(id string) (*model.File, error) {
	var file = &model.File{}
	if err := DB.Model(&model.File{}).Find(file, id).Error; err != nil {
		return nil, err
	}
	return file, nil
}

func CreateFile(file *model.File) error {
	tx := DB.Begin()
	if err := DB.Model(&model.File{}).Create(file).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func CreateAvatar(avatar *model.Avatar) error {
	tx := DB.Begin()
	if err := tx.Model(&model.Avatar{}).Create(avatar).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 更新头像
func UpdateAvatar(avatar *model.Avatar) error {
	tx := DB.Begin()
	if err := tx.Model(&model.Avatar{}).Updates(avatar).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 通过用户名找到对应的头像
func FindAvatarByUsername(username string) (avatar model.Avatar, err error) {
	if err := DB.Model(&model.Avatar{}).Select("path").Where("username = ?", username).Scan(&avatar).Error; err != nil {
		return model.Avatar{}, err
	}
	return avatar, nil
}

// 增加文件下载次数
func AddFileDownloadTimes(fileId string, num uint) error {
	file, err := FindFileByID(fileId)
	if err != nil {
		return err
	}
	file.DownloadTimes += num
	tx := DB.Begin()
	if err := tx.Save(&file).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 删除文件
func DeleteFileById(username string, id int) error {
	file, err := FindFileByID(fmt.Sprintf("%d", id))
	if err != nil {
		return err
	}
	if file.Username != username {
		return errors.New(ErrorNoPermission)
	}
	tx := DB.Begin()
	if err := tx.Delete(file).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := os.Remove(file.Path); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
