package models

import (
	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	Id       int    `json:"Id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
