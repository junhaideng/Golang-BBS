package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("bbs_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func ReleaseToken(username string) (tokenString string, err error) {
	// 过期时间 7天
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Edgar",
			Subject:   "jwt token for auth",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
