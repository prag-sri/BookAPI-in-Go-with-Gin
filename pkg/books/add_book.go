package books

import (
	"net/http"

	//"net/http" is a standard Go package that provides HTTP client and server implementations.

	"github.com/gin-gonic/gin"
	"github.com/prag-sri/go-gin-api-medium/pkg/common/models"
)

type AddBookRequestBody struct {
	Title       string `json:"title"`
	AuthorId    int    `json:"authorId"`
	Description string `json:"description"`

	//It represents the structure of the request body when adding a book.
}

func (h handler) AddBook(c *gin.Context) {

	body := AddBookRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	book.Title = body.Title
	book.AuthorId = body.AuthorId
	book.Description = body.Description

	if result := h.DB.Create(&book); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &book)
}
