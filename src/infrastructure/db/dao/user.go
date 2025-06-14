package infrastructure

import (
	db "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/adapter"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/entities"
	"gorm.io/gorm"
)

type MySQLUserDao struct {
	db *gorm.DB
}

func NewMySQLUserDao(connection *db.DBConnection) *MySQLUserDao {
	return &MySQLUserDao{db: connection.DB}
}

func (dao *MySQLUserDao) GetAllUsers(page, limit int) ([]entities.User, error) {
	var users []entities.User
	if err := dao.db.Limit(limit).Offset((page - 1) * limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (dao *MySQLUserDao) GetUserById(id int) (entities.User, error) {
	var user entities.User
	if err := dao.db.First(&user, id).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}
