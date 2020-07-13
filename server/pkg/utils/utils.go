package utils

import (
	"crypto/md5"
	"encoding/hex"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Token generates random token
func Token(payload string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(payload), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}
