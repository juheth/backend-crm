package infrastructure

import (
	"log"

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
	offset := (page - 1) * limit
	if err := dao.db.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
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

func (dao *MySQLUserDao) CreateUser(user *entities.User) error {
	return dao.db.Create(user).Error
}

func (dao *MySQLUserDao) UpdateUser(user *entities.User) error {
	return dao.db.Model(&entities.User{}).
		Where("id = ?", user.ID).
		Updates(user).Error
}

func (dao *MySQLUserDao) FindByEmail(email string) (entities.User, error) {
	var user entities.User
	err := dao.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (dao *MySQLUserDao) UpdateUserPasswordByEmail(email, hashedPassword string) error {
	err := dao.db.Model(&entities.User{}).Where("email = ?", email).Update("password", hashedPassword).Error
	if err == nil {
		log.Printf("[INFO] Contraseña actualizada para %s", email)
	}
	return err
}
