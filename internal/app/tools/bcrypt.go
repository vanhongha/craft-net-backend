package tools

import (
	"github.com/samber/lo"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(s string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if !lo.IsNil(err) {
		return "", err
	}
	return string(hashed), nil
}

func ComparePassword(hashed string, normal string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(normal))
}
