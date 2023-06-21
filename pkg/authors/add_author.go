package authors

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prag-sri/go-gin-api-medium/pkg/common/models"
)

type AddAuthorRequestBody struct {
	Name    string `json:"name"`
	EmailID string `json:"emailId"`
	Age     uint   `json:"age"`
	Country string `json:"country"`
}

type AddAuthorResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	EmailID string `json:"emailId"`
	Age     uint   `json:"age"`
	Country string `json:"country"`
}

func (h handler) AddAuthor(c *gin.Context) {
	body := AddAuthorRequestBody{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	author := models.Author{
		Name:    body.Name,
		EmailID: body.EmailID,
		Age:     body.Age,
		Country: body.Country,
	}

	if err := h.DB.Create(&author).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response := AddAuthorResponse{
		ID:      author.ID,
		Name:    author.Name,
		EmailID: author.EmailID,
		Age:     author.Age,
		Country: author.Country,
	}

	c.JSON(http.StatusCreated, response)
}
