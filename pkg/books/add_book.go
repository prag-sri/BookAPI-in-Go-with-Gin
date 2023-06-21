package books

import (
	"net/http"

	//"net/http" is a standard Go package that provides HTTP client and server implementations.

	"github.com/gin-gonic/gin"
	"github.com/prag-sri/go-gin-api-medium/pkg/common/models"
)

type AddBookRequestBody struct {
	Title       string `json:"title"`
	AuthorID    int    `json:"authorId"`
	Description string `json:"description"`

	//It represents the structure of the request body when adding a book.
}

type AddBookResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	AuthorID    uint   `json:"authorId"`
	Description string `json:"description"`
}

func (h handler) AddBook(c *gin.Context) {
	var requestBody AddBookRequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	author := models.Author{}
	if err := h.DB.First(&author, requestBody.AuthorID).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	book := models.Book{
		Title:       requestBody.Title,
		Description: requestBody.Description,
		AuthorID:    author.ID, // Set the AuthorID to match the retrieved author's ID
		Author:      author,    // Set the Author to the retrieved author
	}

	if err := h.DB.Create(&book).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Manually retrieve the books associated with the author
	if err := h.DB.Model(&author).Association("Books").Find(&author.Books); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := AddBookResponse{
		ID:          book.ID,
		Title:       book.Title,
		AuthorID:    book.AuthorID,
		Description: book.Description,
	}

	c.JSON(http.StatusCreated, response)
}
