// internal/routes/routes.go

package routes

import (
	"net/http"
	"skylab-book-chameleon/internal/database"
	"skylab-book-chameleon/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *database.DB) *gin.Engine {
	r := gin.Default()

	bookHandlers := &handlers.BookHandlers{DB: db}

	// Serve static files
	r.Static("/static", "./static")

	// Redirect from root path to /static/index.html
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/static/index.html")
	})

	// API routes
	api := r.Group("/api")

	{
		api.POST("/books", bookHandlers.CreateBook)
		api.GET("/books/:id", bookHandlers.GetBook)
		api.PUT("/books/:id", bookHandlers.UpdateBook)
		api.DELETE("/books/:id", bookHandlers.DeleteBook)
		api.GET("/books", bookHandlers.ListBooks)
	}

	return r
}