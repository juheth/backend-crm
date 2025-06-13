package infraestructure

import (
	"fmt"
	"log"

	config "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConnection struct {
	*gorm.DB
}

func NewDBConnection(cfg *config.Config) *DBConnection {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Dbname,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error to connect database : %s", err)
	}
	return &DBConnection{db}
}
