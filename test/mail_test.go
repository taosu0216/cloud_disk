package test

import (
	"cloud_disk/core/utils"
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"net/textproto"
	"testing"
)

func TestMail(t *testing.T) {
	e := &email.Email{
		To:      []string{"1091397182@qq.com"},
		From:    "Taosuu <taosuu@qq.com>",
		Subject: "验证码发送测试",
		//Text:    []byte("Text Body is, of course, supported!"),
		HTML:    []byte("你的验证码为<h1>132456</h1>"),
		Headers: textproto.MIMEHeader{},
	}
	err := e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", "taosuu@qq.com", "ajionyzgpmgrebcc", "smtp.qq.com"),
		&tls.Config{
			InsecureSkipVerify: true,
			ServerName:         "smtp.qq.com",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
}
func TestRandomCode(t *testing.T) {
	fmt.Println(utils.RandomCodeGenerate())
	return
}
