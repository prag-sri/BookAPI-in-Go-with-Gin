package authors

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/authors")
	routes.POST("/", h.AddAuthor)
	routes.GET("/:id", h.GetBooksOfAuthor)
	routes.GET("/", h.GetAuthors)
	routes.DELETE("/:id", h.DeleteAuthor)
}
