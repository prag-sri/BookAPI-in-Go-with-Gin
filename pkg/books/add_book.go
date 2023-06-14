package books

import (
	"net/http"

	//"net/http" is a standard Go package that provides HTTP client and server implementations.

	"github.com/gin-gonic/gin"
	"github.com/prag-sri/go-gin-api-medium/pkg/common/models"
)

type AddBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`

	//It represents the structure of the request body when adding a book.
}

func (h handler) AddBook(c *gin.Context) {

	//This code defines a method named "AddBook" associated with the "handler" struct. The method takes a pointer to a gin.Context object, which represents the HTTP request context.

	body := AddBookRequestBody{}

	//This line creates an instance of the "AddBookRequestBody" struct, initializing it with an empty value.

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//This line binds the JSON data from the request body to the "body" variable using BindJSON method. If there's an error in parsing the JSON or binding it to the struct, an error response with a status code of http.StatusBadRequest is returned.

	var book models.Book

	//This line declares a variable named "book" of type "models.Book".

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	if result := h.DB.Create(&book); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	//This line uses the "Create" method of the "DB" field in the "handler" struct to create a new record in the database using the data from the "book" variable. If there's an error during the creation process, an error response with a status code of http.StatusNotFound is returned.

	c.JSON(http.StatusCreated, &book)

	//This line sends a JSON response to the client with the status code http.StatusCreated (201) and the "book" variable as the response body.
}
