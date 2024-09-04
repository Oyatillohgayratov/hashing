package hashing

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

func Hashing(s string) (string, error) {
	if len(s) == 0 {
		return "", errors.New("Input string cannot be empty")
	}

	hasher := sha256.New()
	hasher.Write([]byte(s))
	hashed := hasher.Sum(nil)

	return hex.EncodeToString(hashed), nil
}

func CheckHash(hash, s string) bool {
	expectedHash, err := Hashing(s)
	if err != nil {
		return false
	}
	return expectedHash == hash
}