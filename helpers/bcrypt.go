package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(password string) string {
	salt := 8
	pass := []byte(password)
	hash, _ := bcrypt.GenerateFromPassword(pass, salt)

	return string(hash)
}

func ComparePass(hashing, password []byte) bool {
	hash, pass := []byte(hashing), []byte(password)
	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}
