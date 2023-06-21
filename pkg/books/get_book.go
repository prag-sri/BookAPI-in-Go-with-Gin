//Here, we just respond with only 1 book based on the ID which we get from a parameter.

package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prag-sri/go-gin-api-medium/pkg/common/models"
)

type GetBookResponseBody struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	AuthorID    uint   `json:"authorId"`
	AuthorName  string
	Description string `json:"description"`
}

func (h handler) GetBook(c *gin.Context) {
	bookID := c.Param("id")

	// Retrieve the book from the database using the bookID
	book := models.Book{}
	if err := h.DB.First(&book, bookID).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// Retrieve the author associated with the book
	author := models.Author{}
	if err := h.DB.First(&author, book.AuthorID).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// Build the response body
	responseBody := GetBookResponseBody{
		ID:          book.ID,
		Title:       book.Title,
		AuthorID:    author.ID,
		AuthorName:  author.Name,
		Description: book.Description,
	}

	c.JSON(http.StatusOK, responseBody)
}
