//Here, we just respond with only 1 book based on the ID which we get from a parameter.

package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetBookResponseBody struct {
	Title       string `json:"title"`
	AuthorId    int    `json:"authorId"`
	AuthorName  string
	Description string `json:"description"`
}

func (h handler) GetBook(c *gin.Context) {

	id := c.Param("id")
	book := GetBookResponseBody{}

	if result := h.DB.Table("books").
		Select("books.title, books.author_id, authors.name AS author_name, books.description").
		Joins("JOIN authors on books.author_id= authors.id").
		Where("books.id =?", id).
		First(&book); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return

		// if result := h.DB.Table("books"): This line selects the "books" table from the database using h.DB.Table("books")

		//Select("books.title, books.author_id, authors.name AS author_name, books.description"): This line specifies the columns to be selected from the "books" table and the "authors" table.

		//.Joins("JOIN authors on books.author_id= authors.id"): This line performs a join operation between the "books" table and the "authors" table based on the author_id column in the "books" table and the id column in the "authors" table.

		//.Where("books.id =?", id): This line adds a condition to filter the rows in the "books" table where the id column matches the provided id variable.

		//.First(&book): This line executes the query and retrieves the first result from the database based on the conditions specified. The result is stored in the book variable.

	}

	c.JSON(http.StatusOK, &book)
}
