package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), 5)

	return string(bytes)
}

func CheckPasswordHash(pass string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))

	return err == nil
}
