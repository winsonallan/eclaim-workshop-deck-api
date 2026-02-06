package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateRandomPassword(length int) (string, string, error) {
	b := make([]byte, length)
	rand.Read(b)

	finStr := base64.URLEncoding.EncodeToString(b)[:length]

	hashed, err := bcrypt.GenerateFromPassword([]byte(finStr), bcrypt.DefaultCost)
	if err != nil {
		return finStr, "", fmt.Errorf("failed to hash password: %v", err)
	}
	return finStr, string(hashed), nil
}
