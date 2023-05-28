package utils

import (
	"fmt"
	"golab/models"

	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB function: Make database connection
func Connect() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(
		&models.User{})

	fmt.Println("Successfully connected!", db)
}
