//Here, we just respond with only 1 book based on the ID which we get from a parameter.

package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prag-sri/go-gin-api-medium/pkg/common/models"
)

func (h handler) GetBook(c *gin.Context) {

	//This code defines a method named "GetBook" associated with the "handler" struct. The method takes a pointer to a gin.Context object, which represents the HTTP request context.

	id := c.Param("id")

	//This line retrieves the value of the "id" parameter from the request URL. It uses the Param method of the gin.Context object to extract the value.

	var book models.Book

	//This line declares a variable named "book" of type "models.Book".

	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	//This line uses the "First" method of the "DB" field in the "handler" struct to retrieve the first book record from the database that matches the given "id". The retrieved book record is stored in the "book" variable. If there's an error during the retrieval process, an error response with a status code of http.StatusNotFound is returned.

	c.JSON(http.StatusOK, &book)

	//This line sends a JSON response to the client with the status code http.StatusOK (200) and the "book" variable as the response body.
}
