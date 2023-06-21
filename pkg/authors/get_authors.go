package authors

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prag-sri/go-gin-api-medium/pkg/common/models"
)

type GetAuthorsResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	EmailID string `json:"emailId"`
	Age     uint   `json:"age"`
	Country string `json:"country"`
}

func (h handler) GetAuthors(c *gin.Context) {

	var authors []models.Author

	if err := h.DB.Find(&authors).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var response []GetAuthorsResponse
	for _, author := range authors {
		response = append(response, GetAuthorsResponse{
			ID:      author.ID,
			Name:    author.Name,
			EmailID: author.EmailID,
			Age:     author.Age,
			Country: author.Country,
		})
	}

	c.JSON(http.StatusOK, response)
}
