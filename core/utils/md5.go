package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5(s string) string {
	token := fmt.Sprintf("%x", md5.Sum([]byte(s)))
	return token
}
