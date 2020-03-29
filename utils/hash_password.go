package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes password with salts
func HashPassword(password string) (string, error) {
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}
