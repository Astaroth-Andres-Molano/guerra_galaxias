// db/database.go

package db

import (
	"fmt"
	"guerra_galaxias/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	dsn := "host=localhost user=postgres password=postgres dbname=starwars port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error al conectar con la base de datos: " + err.Error())
	}

	DB = db
	fmt.Println("conexión a bd exitosa")
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
