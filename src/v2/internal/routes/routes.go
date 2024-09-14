// internal/routes/routes.go

package routes

import (
	"skylab-book-chameleon/internal/database"
	"skylab-book-chameleon/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *database.DB) *gin.Engine {
	r := gin.Default()

	bookHandlers := &handlers.BookHandlers{DB: db}

	r.POST("/books", bookHandlers.CreateBook)
	r.GET("/books/:id", bookHandlers.GetBook)
	r.PUT("/books/:id", bookHandlers.UpdateBook)
	r.DELETE("/books/:id", bookHandlers.DeleteBook)
	r.GET("/books", bookHandlers.ListBooks)

	return r
}
