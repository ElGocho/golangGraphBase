package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(cadena string) (string,error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(cadena), bcrypt.DefaultCost)

	return string(bytes), err
}


func CompareHash(hash, cadena string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(cadena))
}
