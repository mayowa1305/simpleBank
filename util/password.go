package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//HashPassword rerurns the bcrypt hash of a password
func HashPassword(password string) (string, error) {
	HashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password:%v", err)
	}
	return string(HashPassword), nil
}

func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
