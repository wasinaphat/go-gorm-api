package utils_token

import (

	// "github.com/dgrijalva/jwt-go"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-gorm-api/models"
)

type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

func GenerateToken(Id int, Username string) string {

	claims := &models.JWT{
		Id,
		Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    Username,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("dummy"))
	if err != nil {
		panic(err)
	}
	return t
}
func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}

		return []byte("dummy"), nil
	})
}
