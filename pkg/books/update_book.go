package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prag-sri/go-gin-api-medium/pkg/common/models"
)

type UpdateBookRequestBody struct {
	Title       string `json:"title"`
	AuthorId    int    `json:"authorId"`
	Description string `json:"description"`

	//This code defines a new struct type called "UpdateBookRequestBody".
}

func (h handler) UpdateBook(c *gin.Context) {

	//This code defines a method named "UpdateBook" associated with the "handler" struct. The method takes a pointer to a gin.Context object, which represents the HTTP request context.

	id := c.Param("id")

	//This line retrieves the value of the "id" parameter from the request URL.

	body := UpdateBookRequestBody{}

	//This line creates an instance of the "UpdateBookRequestBody" struct, initializing it with an empty value.

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	//This line declares a variable named "book" of type "models.Book"

	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
		//This line uses the "First" method of the "DB" field in the "handler" struct to retrieve the first book record from the database that matches the given "id". The retrieved book record is stored in the "book" variable. If there's an error during the retrieval process, an error response with a status code of http.StatusNotFound is returned.
	}

	book.Title = body.Title
	book.Description = body.Description
	book.AuthorId = body.AuthorId

	var author models.Author
	if result := h.DB.First(&author, book.AuthorId); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	book.Author = author

	h.DB.Save(&book)

	//This line uses the "Save" method of the "DB" field in the "handler" struct to save the changes made to the "book" variable back to the database.

	c.JSON(http.StatusOK, &book)
}
