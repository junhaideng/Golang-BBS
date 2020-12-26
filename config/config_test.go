package config

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestReadConfig(t *testing.T) {
	fmt.Println(viper.GetString("database.mysql.username"))
	fmt.Println(viper.GetString("database.mysql.password"))
	fmt.Println(viper.GetString("database.mysql.password"))

	fmt.Println(viper.GetString("log.filename"))
	fmt.Println(viper.GetString("log.path"))

	fmt.Println(viper.GetString("filepath.base"))
	fmt.Println(viper.GetString("filepath.detail.avatar"))
	fmt.Println(viper.GetString("filepath.detail.carousel"))
	fmt.Println(viper.GetString("filepath.detail.download"))
}
