package authors

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetBookResponseBody struct {
	Title       string `json:"title"`
	AuthorId    int    `json:"authorId"`
	AuthorName  string
	Description string `json:"description"`
}

func (h handler) GetBooksOfAuthor(c *gin.Context) {
	authorId := c.Param("id")

	author_id, err := strconv.Atoi(authorId)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	books := []GetBookResponseBody{}

	query := `SELECT * FROM books WHERE author_id= ?`

	if result := h.DB.Raw(query, author_id).Scan(&books); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	var authorName string
	query = `SELECT name FROM authors WHERE id=?`

	if result := h.DB.Raw(query, author_id).Scan(&authorName); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	for i := range books {
		books[i].AuthorName = authorName
	}

	c.JSON(http.StatusOK, &books)
}
