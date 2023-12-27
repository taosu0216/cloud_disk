package utils

import (
	"cloud_disk/core/pkg"
	"crypto/tls"
	uuid2 "github.com/google/uuid"
	"github.com/jordan-wright/email"
	"log"
	"math/rand"
	"net/smtp"
	"net/textproto"
	"time"
)

func MailCodeSend(emailTo, code string) error {
	to := []string{emailTo}
	e := &email.Email{
		To:      to,
		From:    "Taosuu <taosuu@qq.com>",
		Subject: "验证码发送测试",
		//Text:    []byte("Text Body is, of course, supported!"),
		HTML:    []byte("你的验证码为<h1>" + code + "</h1>"),
		Headers: textproto.MIMEHeader{},
	}
	err := e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", "taosuu@qq.com", "ajionyzgpmgrebcc", "smtp.qq.com"),
		&tls.Config{
			InsecureSkipVerify: true,
			ServerName:         "smtp.qq.com",
		},
	)
	if err != nil {
		return err
	}
	return nil
}
func RandomCodeGenerate() string {
	s := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < pkg.EmailVerificationCodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}
func GenerateUUID() string {
	uuid, err := uuid2.NewUUID()
	if err != nil {
		log.Fatalln("uuid generate fail: ", err)
		return ""
	}

	return uuid.String()
}
