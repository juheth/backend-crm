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
	return nil
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
