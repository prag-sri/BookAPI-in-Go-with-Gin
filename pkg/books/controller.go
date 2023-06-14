package books

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB

	//This block of code defines a new struct type called "handler". The struct has a single field named "DB", which is a pointer to a gorm.DB object. This means that the "handler" struct will have access to the functionality provided by the GORM library for interacting with databases.
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	//This code defines a function named RegisterRoutes that takes two arguments: a pointer to a gin.Engine object (r) and a pointer to a gorm.DB object (db). The gin.Engine object represents the Gin router, and the gorm.DB object represents the database connection.

	h := &handler{
		DB: db,

		//This line creates an instance of the handler struct and assigns it to the variable h. The handler struct is defined elsewhere in the code and has a field DB of type gorm.DB. The DB field is initialized with the provided db argument, which represents the database connection.
	}

	routes := r.Group("/books")
	routes.POST("/", h.AddBook)
	routes.GET("/", h.GetBooks)
	routes.GET("/:id", h.GetBook)
	routes.PUT("/:id", h.UpdateBook)
	routes.DELETE("/:id", h.DeleteBook)
}
