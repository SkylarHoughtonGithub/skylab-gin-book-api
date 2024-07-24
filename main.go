package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a default Gin router
	r := gin.Default()

	// Load HTML templates from the "templates" directory
	r.LoadHTMLGlob("templates/*")

	// Define a route for the root URL
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Define a route for the /button URL to handle AJAX requests
	r.POST("/button", func(c *gin.Context) {
		// Simulate some processing time
		// time.Sleep(2 * time.Second)
		// Handle the button click here
		c.String(http.StatusOK, "Button was clicked!")
	})

	// Start the server on port 8080
	r.Run(":8080")
}
