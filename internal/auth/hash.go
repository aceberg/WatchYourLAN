package auth

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword - generate hash from password
func HashPassword(pw string) string {

	hashed, err := bcrypt.GenerateFromPassword([]byte(pw), 10)
	if err != nil {
		log.Println("PASSWORD ERROR:", err)
		return ""
	}

	return string(hashed)
}

// MatchPasswords - check if password and hash matches
func MatchPasswords(hash, pw string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))

	return err == nil
}
