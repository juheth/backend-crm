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
