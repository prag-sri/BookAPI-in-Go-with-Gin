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

	//This function, Init, is responsible for initializing the database connection and performing any necessary migrations.
	//*gorm.DB represents a pointer to the DB type defined in the gorm package.
	//This block of code opens a connection to the PostgreSQL database using the provided url and returns a *gorm.DB object.

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	//code that opens the database connection using the PostgreSQL driver. It uses the url parameter and an empty gorm.Config{} object.

	if err != nil {
		log.Fatalln(err)
	}

	//If there is an error while opening the database connection, it is logged using log.Fatalln(err). This will print the error message and exit the program.

	db.AutoMigrate(&models.Author{})
	db.AutoMigrate(&models.Book{})

	// This AutoMigrate function will create the books table for us as soon as we run this application.

	return db
}
