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

	if result := h.DB.Preload("Author").First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	//Preload("Author"): This is a method provided by GORM that allows you to preload associations or relationships while fetching records from the database. In this case, Preload("Author") specifies that you want to preload the Author association for the book object. It ensures that when the book is retrieved from the database, the associated Author is also fetched.

	//It fetches the record with the given bookID from the books table and assigns it to the book variable.

	c.JSON(http.StatusOK, &book)

	//This line sends a JSON response to the client with the status code http.StatusOK (200) and the "book" variable as the response body.
}
