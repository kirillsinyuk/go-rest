package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	bcryptHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(bcryptHash)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
