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

func (dao *MySQLClientDao) GetAllClients() ([]*entities.Client, error) {
	var clients []*entities.Client
	err := dao.db.Find(&clients).Error
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func (dao *MySQLClientDao) GetClientByID(id int) (*entities.Client, error) {
	var client entities.Client
	err := dao.db.First(&client, id).Error
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (dao *MySQLClientDao) UpdateClient(client *entities.Client) error {
	return dao.db.Model(&entities.Client{}).
		Where("id = ?", client.ID).
		Updates(map[string]interface{}{
			"name":  client.Name,
			"email": client.Email,
			"phone": client.Phone,
		}).Error
}
