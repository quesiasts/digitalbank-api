package database

import (
	"log"

	"quesiasts/digitalbank-api.git/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	stringConnection := "host=localhost user=postgres password=root dbname=digitalbank port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConnection))
	if err != nil {
		log.Panic("Error to conect with database")
	}
	DB.AutoMigrate(&models.Account{}, &models.Login{}, &models.Transfers{})
}
