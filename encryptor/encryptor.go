package encryptor

import (
	"crypto/rand"
	"crypto/sha512"
	"fmt"
)

// Returns random token with length = 2 * arg
func CreateToken(halfSize int) string {
	token := make([]byte, 64)
	rand.Read(token)
	return fmt.Sprintf("%x", token)
}

// Returns sha512 hash (length of 128) of given string
func GetHash(str string) string {
	h := sha512.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}