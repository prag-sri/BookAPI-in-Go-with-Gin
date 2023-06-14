package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prag-sri/go-gin-api-medium/pkg/common/models"
)

func (h handler) GetBooks(c *gin.Context) {

	//This code defines a method named "GetBooks" associated with the "handler" struct. The method takes a pointer to a gin.Context object, which represents the HTTP request context.

	var books []models.Book

	//This line declares a variable named "books" as a slice of "models.Book".

	if result := h.DB.Find(&books); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	//This line uses the "Find" method of the "DB" field in the "handler" struct to retrieve all book records from the database and store them in the "books" variable. If there's an error during the retrieval process, an error response with a status code of http.StatusNotFound is returned.

	c.JSON(http.StatusOK, &books)

	//This line sends a JSON response to the client with the status code http.StatusOK (200) and the "books" variable as the response body.
}
