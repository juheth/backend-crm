package utils

import (
	"errors"

	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/dto"
)

func ValidateCreateProduct(input dto.CreateProductRequest) error {
	if input.Name == "" {
		return errors.New("el nombre es obligatorio")
	}
	if input.Description == "" {
		return errors.New("la descripción es obligatoria")
	}
	if input.Price <= 0 {
		return errors.New("el precio debe ser mayor que cero")
	}
	if input.Stock < 0 {
		return errors.New("el stock no puede ser negativo")
	}
	if input.Status == "" {
		return errors.New("el estado es obligatorio (debe ser 'activo' o 'inactivo')")
	}
	if input.Status != "activo" && input.Status != "inactivo" {
		return errors.New("el estado debe ser 'activo' o 'inactivo'")
	}
	return nil
}

func ParseStatus(input string) (bool, error) {
	switch input {
	case "activo":
		return true, nil
	case "inactivo":
		return false, nil
	default:
		return false, errors.New("el estado debe ser 'activo' o 'inactivo'")
	}
}

func ValidateUpdateProduct(input dto.UpdateProductRequest) error {
	if input.Name == "" {
		return errors.New("el nombre es obligatorio")
	}
	if input.Description == "" {
		return errors.New("la descripción es obligatoria")
	}
	if input.Price <= 0 {
		return errors.New("el precio debe ser mayor que cero")
	}
	if input.Stock < 0 {
		return errors.New("el stock no puede ser negativo")
	}
	return nil
}
