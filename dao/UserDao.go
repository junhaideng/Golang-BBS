package dao

import (
	"bbs/model"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

var (
	ErrUserAlreadyExist = errors.New("已存在该用户")
	ErrUserNotExist     = errors.New("用户不存在")
)

// 更具用户名查找用户，返回User类型的引用
func FindUserByUsername(username string) (*model.User, error) {
	var user = &model.User{}
	if err := DB.Model(&model.User{}).Find(user, "username = ?", username).Error; err != nil {
		// 如果没有查找到对应的用户
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserNotExist
		}
		return nil, err
	}
	return user, nil
}

// 创建用户，同时根据用户名会查询用户是否存在
// error -> 创建的时候产生的问题
func CreateUser(user *model.User) error {
	tx := DB.Begin()
	var temp model.User
	tx.Model(&model.User{}).Where("username = ?", user.Username).Find(&temp)
	// 如果找到对应的用户， 说明用户已经存在，不能创建
	fmt.Printf("temp: %#v", temp)
	if len(temp.Username) > 0 {
		return ErrUserAlreadyExist
	}

	if err := tx.Model(&model.User{}).Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

type UserInfo struct {
	Academy     string `json:"academy" form:"academy"`
	Email       string `json:"email" form:"email"`
	Grade       string `json:"grade" form:"grade"`
	Age         uint8  `json:"age" form:"age"`
	Gender      string `json:"gender" form:"gender"`
	Description string `json:"description" form:"description"`
}

// 更新用户的信息
func UpdateUser(user model.User, userInfo UserInfo) error {
	tx := DB.Begin()
	fmt.Printf("%#v", userInfo)
	if err := tx.Model(&user).Updates(userInfo).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

type Password struct {
	Password string
}

func GetPasswordByUsername(username string) (string, error) {
	var password Password
	if err := DB.Model(&model.User{}).Select("password").Where("username = ?", username).Scan(&password).Error; err != nil {
		return "", err
	}
	return password.Password, nil
}

func UpdatePasswordByUsername(username string, password string) error {
	tx := DB.Begin()
	if err := tx.Model(&model.User{}).Where("username = ?", username).Update("password", password).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func DeleteAccountByUsername(username string) error {
	tx := DB.Begin()
	if err := tx.Model(&model.User{}).Where("username = ?", username).Delete(model.User{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
