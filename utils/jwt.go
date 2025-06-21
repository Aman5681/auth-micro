package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(emailId string, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"emailId": emailId,
		"userId":  userId,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
