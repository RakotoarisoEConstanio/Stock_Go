package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"symrise/models"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	dsn := "user=postgres password=password dbname=GestionStock host=localhost port=5432 sslmode=disable"

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("impossible de se connecter")
	}

	DB = connection
	connection.AutoMigrate(
		&models.BonEntree{},
		&models.BonSortie{},
		&models.Materiaux{},
		&models.BonEntree{},
	)

	return connection, nil
}
