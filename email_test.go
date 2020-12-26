package main

import (
	"bbs/utils"
	"github.com/spf13/viper"
	"testing"
)

func TestEmail(t *testing.T) {
	if err := utils.SendEmail(viper.GetStringSlice("email.target"), "使用golang发送", utils.TextType, "hhh"); err != nil {
		t.Fatal("send email error: ", err.Error())
		return
	}
	t.Log("Pass")
}
