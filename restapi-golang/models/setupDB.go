package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB maintain connection to the database
var DB *gorm.DB

// ConnectDatabase opens a connection to the SQLite database
func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Task{}, &User{})
	if err != nil {
		return
	}
	// Store the connection to the database in a global variable DB for use anywhere in the application.
	DB = database
}
