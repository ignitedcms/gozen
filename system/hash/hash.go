/*
|---------------------------------------------------------------
| Hash utility helper
|---------------------------------------------------------------
|
| Helper for password management
| and generating random keys for sessions/tokens
|
| @license: MIT
| @version: 1.0
| @since: 1.0
*/
package hash

import (
	"crypto/rand"
	"encoding/hex"
	//"fmt"
	"golang.org/x/crypto/bcrypt"
	//"log"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// generate 32 byte key for sessions and csrf
func GenerateKey() (string, error) {
	key := make([]byte, 32) // 32 bytes = 256 bits
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}
