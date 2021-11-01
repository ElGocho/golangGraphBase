package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(cadena string) (string,error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(cadena), bcrypt.DefaultCost)

	return string(bytes), err
}
