package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prag-sri/go-gin-api-medium/pkg/common/models"
)

type GetBooksResponseBody struct {
	ID          uint              `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Author      GetAuthorResponse `json:"author"`
}

type GetAuthorResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (h handler) GetBooks(c *gin.Context) {
	var books []models.Book

	// Query books table
	if err := h.DB.Find(&books).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Get the author IDs from the books
	var authorIDs []uint
	for _, book := range books {
		authorIDs = append(authorIDs, book.AuthorID)
	}

	// Query authors table using the author IDs
	var authors []models.Author
	if err := h.DB.Where(authorIDs).Find(&authors).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Create a map of author ID to author name for easy lookup
	authorMap := make(map[uint]string)
	for _, author := range authors {
		authorMap[author.ID] = author.Name
	}

	// Build the response body
	var responseBody []GetBooksResponseBody
	for _, book := range books {
		responseBody = append(responseBody, GetBooksResponseBody{
			ID:          book.ID,
			Title:       book.Title,
			Description: book.Description,
			Author: GetAuthorResponse{
				ID:   book.AuthorID,
				Name: authorMap[book.AuthorID], // Use the author's name from the map
			},
		})
	}

	c.JSON(http.StatusOK, responseBody)
}
