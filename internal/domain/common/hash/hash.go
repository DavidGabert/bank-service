package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func Verify(password, hashedPassword string) bool {
	return Hash(password) == hashedPassword
}
