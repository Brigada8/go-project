package repositories

import (
	"fmt"
	"golab/internal"

	"github.com/glebarez/sqlite"

	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB function: Make database connection
func Connect() {
	var err error
	connection, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = connection

	// Migrate the schema
	DB.AutoMigrate(
		&internal.User{})

	fmt.Println("Successfully connected!", DB)
}
