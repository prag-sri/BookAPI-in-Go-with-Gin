package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prag-sri/go-gin-api-medium/pkg/common/models"
)

func (h handler) DeleteBook(c *gin.Context) {
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

	// Remove the book from the author's list of books
	for i, b := range author.Books {
		if b.ID == book.ID {
			author.Books = append(author.Books[:i], author.Books[i+1:]...)
			break
		}
	}

	// Delete the book from the database
	if err := h.DB.Delete(&book).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Save the changes to the author
	if err := h.DB.Save(&author).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
