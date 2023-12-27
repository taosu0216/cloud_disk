package utils

import (
	"cloud_disk/core/pkg"
	"github.com/golang-jwt/jwt"
	"log"
)

func GenerateToken(id uint64, name, identity string) string {
	userclaim := &pkg.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userclaim)
	signedtoken, err := token.SignedString(pkg.GetSecretKey())
	if err != nil {
		log.Fatalln("jwt generated err is :", err)
		return ""
	}
	return signedtoken
}
