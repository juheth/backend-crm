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

func (dao *MySQLProductDao) GetAll() ([]*entities.Product, error) {
	var products []*entities.Product
	err := dao.db.Where("status = ?", true).Find(&products).Error
	return products, err
}

func (dao *MySQLProductDao) GetProductByID(id int) (*entities.Product, error) {
	var product entities.Product
	err := dao.db.Where("id = ? AND status = ?", id, true).First(&product).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}
	return &product, nil
}

func (dao *MySQLProductDao) UpdateProduct(product *entities.Product) error {
	return dao.db.Model(&entities.Product{}).
		Where("id = ? AND status = ?", product.ID, true).
		Updates(map[string]interface{}{
			"name":        product.Name,
			"description": product.Description,
			"price":       product.Price,
			"stock":       product.Stock,
		}).Error
}

func (dao *MySQLProductDao) DeactivateProduct(id int) error {
	return dao.db.Model(&entities.Product{}).
		Where("id = ? AND status = ?", id, true).
		Update("status", false).Error
}

func (dao *MySQLProductDao) ActivateProduct(id int) error {
	return dao.db.Model(&entities.Product{}).
		Where("id = ? AND status = ?", id, false).
		Update("status", true).Error
}

func (dao *MySQLProductDao) GetAllDeactivated() ([]*entities.Product, error) {
	var products []*entities.Product
	err := dao.db.Where("status = ?", false).Find(&products).Error
	return products, err
}
func (dao *MySQLProductDao) GetProductByIDAnyStatus(id int) (*entities.Product, error) {
	var product entities.Product
	err := dao.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (dao *MySQLProductDao) DeleteProduct(id int) error {
	return dao.db.Where("id = ?", id).Delete(&entities.Product{}).Error
}
