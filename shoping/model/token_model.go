package model

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	User User
	Type     string //"REFRESH_TOKEN"表示为一个refresh token，"TOKEN"表示为一个token
	Time     time.Time
	jwt.StandardClaims
}

