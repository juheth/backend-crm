package utils

import (
	"errors"
	"fmt"

	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/domain/dto"
)

func ValidateCreateOrder(req dto.CreateOrderRequest, clientDao *dao.MySQLClientDao) error {
	if req.ClientID <= 0 {
		return errors.New("El cliente es obligatorio")
	}
	exists, err := clientDao.ExistsByID(req.ClientID)
	if err != nil {
		return fmt.Errorf("Error al validar cliente: %v", err)
	}
	if !exists {
		return errors.New("El cliente no existe")
	}

	if req.PaymentMethodID <= 0 {
		return errors.New("El método de pago es obligatorio")
	}
	if len(req.Items) == 0 {
		return errors.New("Debe agregar al menos un producto al pedido")
	}
	for i, item := range req.Items {
		if item.ProductID <= 0 {
			return fmt.Errorf("El producto en la posición %d es inválido", i+1)
		}
		if item.Quantity <= 0 {
			return fmt.Errorf("La cantidad para el producto en la posición %d debe ser mayor a cero", i+1)
		}
	}
	return nil
}
