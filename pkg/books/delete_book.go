package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prag-sri/go-gin-api-medium/pkg/common/models"
)

func (h handler) DeleteBook(c *gin.Context) {

	//This code defines a method named "DeleteBook" associated with the "handler" struct. The method takes a pointer to a gin.Context object, which represents the HTTP request context.

	id := c.Param("id")

	//This line retrieves the value of the "id" parameter from the request URL.

	var book models.Book

	//This line declares a variable named "book" of type "models.Book". It likely represents a book model defined in the imported package.

	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return

		//This line uses the "First" method of the "DB" field in the "handler" struct to retrieve the first book record from the database that matches the given "id". The retrieved book record is stored in the "book" variable. If there's an error during the retrieval process, an error response with a status code of http.StatusNotFound is returned.
	}

	h.DB.Delete(&book)

	//This line uses the "Delete" method of the "DB" field in the "handler" struct to delete the book record from the database.

	c.Status(http.StatusOK)

	// This line sets the response status code to http.StatusOK (200). It indicates that the book deletion was successful. No response body is provided in this case.
}
