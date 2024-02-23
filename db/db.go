package db

import (
	"github.com/JohnKucharsky/echo_gorm/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type ApiConfig struct {
	DB *gorm.DB
}

func DatabaseConnection(dbAddressString string) *ApiConfig {
	db, err := gorm.Open(
		postgres.Open(dbAddressString), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		log.Fatal("Can't connect to db", err.Error())
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
	); err != nil {
		log.Fatal("Can't migrate", err.Error())
	}

	apiCfg := ApiConfig{DB: db}

	return &apiCfg
}
