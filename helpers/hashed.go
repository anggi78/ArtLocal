package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	result, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(result)
}

func ComparePassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

// func HassPass(pass string) (string, error) {
// 	p := []byte(pass)
// 	salt := 8
// 	Hass, err := bcrypt.GenerateFromPassword(p, salt)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(Hass), nil
// }

// func ComparePass(hass, pass []byte) (bool, error) {
// 	err := bcrypt.CompareHashAndPassword(hass, pass)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }