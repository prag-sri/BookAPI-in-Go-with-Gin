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

	id := c.Param("id")

	body := UpdateBookRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	book.Title = body.Title
	book.Description = body.Description
	book.AuthorId = body.AuthorId

	h.DB.Save(&book)
	c.JSON(http.StatusOK, &book)
}
