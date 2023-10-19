package helpers

import "golang.org/x/crypto/bcrypt"

func Hash(pass string) (string, error) {
	password := []byte(pass)
	value := 8
	HashPass, err := bcrypt.GenerateFromPassword(password, value)
	if err != nil {
		return "", err
	}
	return string(HashPass), nil
}

func Compare(hash, pass []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, pass)
	if err != nil {
		return false, err
	}
	return true, nil
}