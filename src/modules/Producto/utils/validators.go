package utils

import (
	"errors"
	"strings"
)

func ValidateProductInput(name, description string, price float64, stock int) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("El nombre es obligatorio")
	}
	if len(name) < 3 {
		return errors.New("El nombre debe tener al menos 3 caracteres")
	}
	if strings.TrimSpace(description) == "" {
		return errors.New("La descripción es obligatoria")
	}
	if price <= 0 {
		return errors.New("El precio debe ser mayor a 0")
	}
	if stock < 0 {
		return errors.New("El stock no puede ser negativo")
	}
	return nil
}
