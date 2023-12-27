package pkg

import (
	"github.com/golang-jwt/jwt"
)

var jwtSecretKey = []byte("1433223_taosu_Cloud_disk")

type UserClaim struct {
	Id       uint64
	Identity string
	Name     string
	jwt.StandardClaims
}

func GetSecretKey() []byte {
	return jwtSecretKey
}
