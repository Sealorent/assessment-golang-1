package utils

import "golang.org/x/crypto/bcrypt"

func Hash(data []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(data, 8)
}

func HashMatched(hash []byte, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, plain)
	return err == nil
}
