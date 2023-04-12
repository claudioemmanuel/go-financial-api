package utils

import (
	"strconv"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func GenerateFromString(s string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
}

func CompareHashAndString(s string, hash []byte) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(s))
}

func ConvertStringToUint(s string) uint {
	u, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0
	}
	return uint(u)
}
