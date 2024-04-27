package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(value string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CompareHash(hash string, value string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))

	return err == nil
}
