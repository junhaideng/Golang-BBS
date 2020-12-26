package log

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// 初始化日志
func init() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(filepath.Join(cwd, viper.GetString("log.path"), viper.GetString("log.filename")), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)
}
