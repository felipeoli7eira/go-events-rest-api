package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const JWT_SECRET_KEY = "supersecret"

func GenerateJWT(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(JWT_SECRET_KEY))
}
