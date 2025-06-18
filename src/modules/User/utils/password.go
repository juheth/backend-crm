package utils

import (
	"fmt"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidatePassword(password string) error {
	var hasMinLen, hasUpper, hasNumber, hasSymbol bool

	if len(password) >= 8 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSymbol = true
		}
	}

	if !hasMinLen {
		return fmt.Errorf("la contraseña debe tener al menos 8 caracteres")
	}
	if !hasUpper {
		return fmt.Errorf("la contraseña debe contener al menos una letra mayúscula")
	}
	if !hasNumber {
		return fmt.Errorf("la contraseña debe contener al menos un número")
	}
	if !hasSymbol {
		return fmt.Errorf("la contraseña debe contener al menos un símbolo")
	}

	return nil
}
