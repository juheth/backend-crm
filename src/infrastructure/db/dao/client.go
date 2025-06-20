package infrastructure

import (
	db "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/adapter"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/domain/entities"

	"gorm.io/gorm"
)

type MySQLClientDao struct {
	db *gorm.DB
}

func NewMySQLClientDao(connection *db.DBConnection) *MySQLClientDao {
	return &MySQLClientDao{db: connection.DB}
}

func (dao *MySQLClientDao) CreateClient(client *entities.Client) error {
	return dao.db.Create(client).Error
}

func (dao *MySQLClientDao) ExistsByEmailOrPhone(email, phone string) (bool, error) {
	var count int64
	err := dao.db.Model(&entities.Client{}).
		Where("email = ? OR phone = ?", email, phone).
		Count(&count).Error
	return count > 0, err
}
