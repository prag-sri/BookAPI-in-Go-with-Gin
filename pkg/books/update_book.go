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

type UpdateBookResponseBody struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	AuthorID    uint   `json:"authorId"`
	Description string `json:"description"`
}

func (h handler) UpdateBook(c *gin.Context) {
	bookID := c.Param("id")

	var requestBody UpdateBookRequestBody

	// Bind the JSON request body to the UpdateBookRequestBody struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Retrieve the book from the database using the bookID
	book := models.Book{}
	if err := h.DB.First(&book, bookID).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// Retrieve the author from the database using the authorID
	author := models.Author{}
	if err := h.DB.First(&author, requestBody.AuthorId).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// Update the book's properties
	book.Title = requestBody.Title
	book.Description = requestBody.Description
	book.Author = author

	// Save the updated book to the database
	if err := h.DB.Save(&book).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := UpdateBookResponseBody{
		ID:          book.ID,
		Title:       book.Title,
		AuthorID:    book.AuthorID,
		Description: book.Description,
	}

	c.JSON(http.StatusOK, &response)
}
