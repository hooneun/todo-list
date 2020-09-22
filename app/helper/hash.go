package helper

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func makeHash(s *string) error {
	if s == nil {
		return errors.New("Reference provided for hashing string is nil")
	}

	sBytes := []byte(*s)
	hashed, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	*s = string(hashed[:])
	return nil
}

func checkHash(hash, s string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(s)) == nil
}
