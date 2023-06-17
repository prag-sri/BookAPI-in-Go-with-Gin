package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetBooks(c *gin.Context) {

	books := []GetBookResponseBody{}

	err := h.DB.Table("books").
		Select("books.title, books.author_id, authors.name AS author_name, books.description").
		Joins("JOIN authors ON books.author_id = authors.id").
		Find(&books).Error

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, books)
}
