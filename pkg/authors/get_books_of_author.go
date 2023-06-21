package authors

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetBookResponseBody struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	AuthorID    uint   `json:"authorId"`
	AuthorName  string
	Description string `json:"description"`
}

func (h handler) GetBooksOfAuthor(c *gin.Context) {
	authorID := c.Param("id")

	// Check if authorID is empty
	if authorID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Author ID is missing or invalid"})
		return
	}

	// Convert authorID to uint
	authorIDUint, err := strconv.ParseUint(authorID, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var books []GetBookResponseBody

	if err := h.DB.Table("books").
		Select("books.id, books.title, books.author_id, authors.name AS author_name, books.description").
		Joins("JOIN authors ON books.author_id = authors.id").
		Where("books.author_id = ?", authorIDUint).
		Scan(&books).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, books)
}
