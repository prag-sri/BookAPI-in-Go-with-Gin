package authors

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prag-sri/go-gin-api-medium/pkg/common/models"
)

type AddAuthorRequestBody struct {
	Name    string `json:"name"`
	EmailId string `json:"emailId"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}

func (h handler) AddAuthor(c *gin.Context) {
	body := AddAuthorRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var author models.Author

	author.Name = body.Name
	author.EmailId = body.EmailId
	author.Age = body.Age
	author.Country = body.Country

	if result := h.DB.Create(&author); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &author)
}
