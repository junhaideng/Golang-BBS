package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func init() {
	base, err := os.Getwd()
	fmt.Println(base)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	setDefault()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	// check configuration
	check()

	// 文件目录
	fp := filepath.Join(base, viper.GetString("filepath.base"))
	initPath(fp, viper.GetString("filepath.detail.avatar"))
	initPath(fp, viper.GetString("filepath.detail.carousel"))
	initPath(fp, viper.GetString("filepath.detail.download"))
}

// 设置默认配置
func setDefault() {
	// 设置默认目录配置
	viper.SetDefault("filepath.base", "file")
	viper.SetDefault("filepath.detail.avatar", "avatar")
	viper.SetDefault("filepath.detail.carousel", "carousel")
	viper.SetDefault("filepath.detail.download", "download")

	// 设置默认日志配置
	viper.SetDefault("log.filename", "bbs.log")
	viper.SetDefault("log.path", ".")
}

// 初始化目录，如果目录不存在，则创建目录
func initPath(base, path string) {
	p := filepath.Join(base, path)
	_, err := os.Stat(p)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(p, 0666); err != nil {
			fmt.Println("创建目录失败")
			return
		}
	}
}

// 检查配置
func check() {
	// baidu api
	if viper.Get("baidu.key") == nil {
		panic("no specified baidu api key, please visit url 'http://lbsyun.baidu.com/index.php' and get key")
	}

	// database
	if viper.Get("database.mysql") == nil {
		panic("no specified database configuration")
	}

	// email
	if viper.Get("email.username") == nil || viper.Get("email.password") == nil {
		panic("no specified email username or password ")
	}
	if len(viper.GetStringSlice("email.target")) == 0 {
		panic("no specified email target to report error")
	}
}
