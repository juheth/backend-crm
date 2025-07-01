package infrastructure

import (
	db "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/adapter"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/domain/entities"
	"gorm.io/gorm"
)

type MySQLOrderDao struct {
	db *gorm.DB
}

func NewMySQLOrderDao(connection *db.DBConnection) *MySQLOrderDao {
	return &MySQLOrderDao{db: connection.DB}
}

func (dao *MySQLOrderDao) CreateOrder(order *entities.Order) error {
	return dao.db.Create(order).Error
}

func (dao *MySQLOrderDao) GetProductPrice(productID int) (float64, error) {
	var price float64
	err := dao.db.Table("products").Select("price").Where("id = ?", productID).Scan(&price).Error
	if err != nil {
		return 0, err
	}
	return price, nil
}

func (dao *MySQLClientDao) ExistsByID(id int) (bool, error) {
	var count int64
	err := dao.db.Table("clients").Where("id = ?", id).Count(&count).Error
	return count > 0, err
}

func (dao *MySQLOrderDao) GetAllOrders() ([]*entities.Order, error) {
	var orders []*entities.Order
	err := dao.db.Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (dao *MySQLOrderDao) GetOrderByID(id int) (*entities.Order, error) {
	var order entities.Order
	err := dao.db.Preload("Items").First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (dao *MySQLOrderDao) UpdateOrderStatus(orderID int, status string) error {
	return dao.db.Model(&entities.Order{}).
		Where("id = ?", orderID).
		Update("status", status).Error
}

func (dao *MySQLOrderDao) DeleteOrder(orderID int) error {
	tx := dao.db.Begin()
	if err := tx.Where("order_id = ?", orderID).Delete(&entities.OrderItem{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&entities.Order{}, orderID).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (dao *MySQLOrderDao) GetOrdersByClient(clientId int) ([]*entities.Order, error) {
	var orders []*entities.Order
	err := dao.db.Preload("Items").Where("client_id = ?", clientId).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (dao *MySQLOrderDao) GetOrdersByStatus(status string) ([]*entities.Order, error) {
	var orders []*entities.Order
	err := dao.db.Preload("Items").Where("status = ?", status).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (dao *MySQLOrderDao) GetOrderSummary() (totalPedidosMes int, ventasTotalesMes float64, ticketPromedio float64, porEstado map[string]int, err error) {
	var totals struct {
		TotalPedidosMes  int
		VentasTotalesMes float64
		TicketPromedio   float64
	}
	err = dao.db.Raw(`
		SELECT
			COUNT(*) AS total_pedidos_mes,
			IFNULL(SUM(total),0) AS ventas_totales_mes,
			IFNULL(AVG(total),0) AS ticket_promedio
		FROM orders
		WHERE MONTH(order_date) = MONTH(CURRENT_DATE())
		  AND YEAR(order_date) = YEAR(CURRENT_DATE())
	`).Scan(&totals).Error
	if err != nil {
		return 0, 0, 0, nil, err
	}

	type EstadoCount struct {
		Status   string
		Cantidad int
	}
	var estados []EstadoCount
	err = dao.db.Raw(`
		SELECT status, COUNT(*) as cantidad
		FROM orders
		WHERE MONTH(order_date) = MONTH(CURRENT_DATE())
		  AND YEAR(order_date) = YEAR(CURRENT_DATE())
		GROUP BY status
	`).Scan(&estados).Error
	if err != nil {
		return 0, 0, 0, nil, err
	}

	porEstado = make(map[string]int)
	for _, e := range estados {
		porEstado[e.Status] = e.Cantidad
	}

	return totals.TotalPedidosMes, totals.VentasTotalesMes, totals.TicketPromedio, porEstado, nil
}
