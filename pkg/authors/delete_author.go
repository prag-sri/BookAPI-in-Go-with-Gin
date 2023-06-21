package authors

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prag-sri/go-gin-api-medium/pkg/common/models"
)

func (h handler) DeleteAuthor(c *gin.Context) {
	authorID := c.Param("id")

	// Delete books associated with the author
	if err := h.DB.Where("author_id = ?", authorID).Delete(&models.Book{}).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Delete the author
	if err := h.DB.Delete(&models.Author{}, authorID).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
