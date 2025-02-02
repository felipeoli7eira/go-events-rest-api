package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPlainText(plainText string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost )

	return string(bytes), err
}

func PlainTextAndHashMatch(hashed, plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plainText))
	return err == nil
}
