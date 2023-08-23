package security

import "golang.org/x/crypto/bcrypt"

func HashGen(str string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
}

func HashCompare(hash, str string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
}
