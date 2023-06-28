// db/database.go

package db

import (
	"guerra_galaxias/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	dsn := "host=localhost user=postgres password=postgres dbname=starwars port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func MigrateDB() error {
	err := DB.AutoMigrate(&models.SatelliteData{})
	if err != nil {
		return err
	}

	// Add migration logic for other models if needed

	return nil
}
