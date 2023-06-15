package authors

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prag-sri/go-gin-api-medium/pkg/common/models"
)

func (h handler) GetBooksOfAuthor(c *gin.Context) {
	authorId := c.Param("id")

	//These lines retrieve the value of the id parameter from the URL path using c.Param("id").

	author_id, err := strconv.Atoi(authorId)

	//strconv.Atoi(authorId) converts the authorId string to an integer and assigns it to the author_id variable.

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var author models.Author

	if result := h.DB.First(&author, author_id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var books []models.Book

	//These lines declare a variable books as a slice of models.Book

	if result := h.DB.Where("author_id = ?", author_id).Preload("Author").Find(&books); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return

		//It uses h.DB.Where("author_id= ?", author_id).Find(&books) to fetch all the book records from the database that match the provided author_id using GORM.
	}

	c.JSON(http.StatusOK, &books)
}
