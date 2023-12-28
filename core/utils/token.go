package utils

import (
	"cloud_disk/core/pkg"
	"errors"
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
func AnalyazeToken(token string) (*pkg.UserClaim, error) {
	uc := new(pkg.UserClaim)
	claim, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return pkg.GetSecretKey(), nil
	})
	if err != nil {
		return nil, err
	}
	if !claim.Valid {
		return nil, errors.New("token戳啦")
	}
	return uc, nil
}
