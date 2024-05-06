package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPssword give the bycrypt hash
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error to hash passowrd %s", password)
	}
	return string(hashedPassword), nil
}

// CheckPassword verify if is a valid password
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
