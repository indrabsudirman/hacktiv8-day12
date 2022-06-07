package helper

import "golang.org/x/crypto/bcrypt"

const salt = 8

func HashPass(p string) (string, error) {
	password := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(password, salt)
	return string(hash), err
}

func CheckPass(dbPass, uPass string) error {
	hash, pass := []byte(dbPass), []byte(uPass)
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err
}
