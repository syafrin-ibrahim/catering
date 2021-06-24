package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}

func CheckPasswordHash(pass, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		return false, err
	}

	return true, nil
}
