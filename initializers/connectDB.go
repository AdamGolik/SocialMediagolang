package initializers

import (
	"1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

// an variable  that we can use outside of this file
var DB *gorm.DB

// func that will let us allow to create or connect to the database
func ConnectDB() {
	var err error
	// Konfiguracja DSN dla PostgreSQL
	dsn := "host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}

// func that will migrate tables in the database
func CreateDb() {
	err := DB.AutoMigrate(&models.User{}, &models.Image{})
	if err != nil {
		panic("Failed to create tables: " + err.Error())
	}
}
