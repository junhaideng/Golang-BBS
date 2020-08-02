package common

import (
	"fmt"
	"os"
	"path/filepath"
)

var ProjectPath string

func init() {
	ProjectPath, err := os.Getwd()
	if err != nil {
		fmt.Println("获取工作路径失败")
		return
	}

	path := filepath.Join(ProjectPath, FileUploadPath)
	_, err = os.Stat(path)
	fmt.Println("Hello, this is the init function of common")
	if os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0666); err != nil {
			fmt.Println("创建文件目录失败")
			return
		}
	} else {
		fmt.Println("目录已存在")
	}

	path = filepath.Join(ProjectPath, AvatarPath)
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0666); err != nil {
			fmt.Println("创建头像目录失败")
			return
		}
	} else {
		fmt.Println("目录已存在")
	}

}

const (
	// 在个人中心保存文件的位置
	// 为当前工作目录下
	FileUploadPath = "/file/download/"

	// 头像保存的位置
	AvatarPath = "/file/avatar/"
)
