package repository

import (
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *entities.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user *entities.User) error {
	return r.db.Create(user).Error
}
