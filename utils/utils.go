package utils

import (
	"fmt"
	"golab/models"

	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

// ConnectDB function: Make database connection
func ConnectDB() *gorm.DB {

	//Define DB connection string
	// dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", databaseHost, username, databaseName, password)

	//connect to db URI
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	// close db when not in use

	// Migrate the schema
	db.AutoMigrate(
		&models.User{})

	fmt.Println("Successfully connected!", db)
	return db
}
