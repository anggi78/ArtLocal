package helpers

import "golang.org/x/crypto/bcrypt"

func Hash(pass string) ([]byte, error) {
	password := []byte(pass)
	cost := 8
	hashedPass, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		return nil, err
	}
	return hashedPass, nil
}

func Compare(hash, pass []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, pass)
	if err != nil {
		return false, err
	}
	return true, nil
}