package db

//This line specifies that the code in this file belongs to the db package.

import (
	"log"

	//This package provides a simple logging interface that is used to print error messages in case of a database connection error.

	"github.com/prag-sri/go-gin-api-medium/pkg/common/models"

	"gorm.io/driver/postgres"

	// This package provides the PostgreSQL database driver for the gorm package.

	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Author{})
	db.AutoMigrate(&models.Book{})

	// Define the association between Author and Book models
	db.Model(&models.Author{}).Association("Books")
	db.Model(&models.Book{}).Association("Author")

	return db
}
