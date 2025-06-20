package utils

import (
	"errors"
	"regexp"
	"strings"
)

func ValidateClientInput(name, email, phone string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("El nombre es obligatorio")
	}
	if !isValidEmail(email) {
		return errors.New("Email inválido")
	}
	if len(phone) < 7 {
		return errors.New("Teléfono inválido")
	}
	return nil
}

func isValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(regex).MatchString(email)
}
