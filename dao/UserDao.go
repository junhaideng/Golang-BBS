package dao

import (
	"bbs/model"
	"errors"
)

var (
	ErrorUserAlreadyExist = errors.New("用户已存在")
	ErrorUserNotExist     = errors.New("用户不存在")
)

// 更具用户名查找用户，返回User类型的引用
func FindUserByUsername(username string) (*model.User, error) {
	var user = &model.User{}
	if err := DB.Model(&model.User{}).Find(user, "username = ?", username).Error; err != nil {
		return nil, ErrorUserNotExist
	}
	return user, nil
}

// 创建用户，同时根据用户名会查询用户是否存在
func CreateUser(user *model.User) error {
	tx := DB.Begin()
	if !tx.Model(&model.User{}).Where("username = ?", user.Username).RecordNotFound() {
		return ErrorUserAlreadyExist
	}

	if err := tx.Model(&model.User{}).Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
