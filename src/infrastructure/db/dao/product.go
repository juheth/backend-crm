package infrastructure

import (
	db "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/adapter"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/entities"
	"gorm.io/gorm"
)

type MySQLProductDao struct {
	db *gorm.DB
}

func NewMySQLProductDao(connection *db.DBConnection) *MySQLProductDao {
	return &MySQLProductDao{db: connection.DB}
}

func (dao *MySQLProductDao) Create(product *entities.Product) error {
	return dao.db.Create(product).Error
}

func (dao *MySQLProductDao) ExistsByName(name string) (bool, error) {
	var count int64
	err := dao.db.Model(&entities.Product{}).Where("name = ?", name).Count(&count).Error
	return count > 0, err
}
