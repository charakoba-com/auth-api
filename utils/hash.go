package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

// HashPassword hashes given string with sha512
func HashPassword(password, salt string) string {
	hash := sha512.New()
	for i := 0; i < 29; i++ {
		hash.Reset()
		hash.Write([]byte(password))
		password = hex.EncodeToString(hash.Sum([]byte(salt)))
	}
	hash.Reset()
	hash.Write([]byte(password))
	password = hex.EncodeToString(hash.Sum(nil))
	return password
}
