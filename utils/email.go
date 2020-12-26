package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"net/smtp"
	"strings"
)

const (
	TextType = "text/plain"
	HtmlType = "text/html"
)

func SendEmail(targets []string, subject, contentType, body string) error {
	username := viper.GetString("email.username")
	password := viper.GetString("email.password")
	server := viper.GetString("email.server")
	to := strings.Join(targets, ",")
	msg := fmt.Sprintf("From: <%s> \r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: %s; charset=UTF-8\r\n\r\n%s",
		username,
		to,
		subject,
		contentType,
		[]byte(body),
	)
	auth := smtp.PlainAuth("", username, password, strings.Split(server, ":")[0])
	if err := smtp.SendMail(viper.GetString("email.server"), auth, username, targets, []byte(msg)); err != nil {
		return err
	}
	return nil
}

func SendEmailToTarget(subject, contentType, body string) error {
	targets := viper.GetStringSlice("email.target")
	return SendEmail(targets, subject, contentType, body)
}
