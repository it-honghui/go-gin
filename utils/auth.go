package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) string {
	const cost = 10
	encryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(encryptPassword)
}
