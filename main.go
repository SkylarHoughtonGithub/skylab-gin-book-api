package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Animal represents data about an animal.
type Animal struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

var animals = []Animal{
	{ID: "1", Name: "Lion", Type: "Mammal"},
	{ID: "2", Name: "Eagle", Type: "Bird"},
}

// Handlers

func getAnimals(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, animals)
}

func getAnimal(c *gin.Context) {
	id := c.Param("id")
	for _, animal := range animals {
		if animal.ID == id {
			c.IndentedJSON(http.StatusOK, animal)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "animal not found"})
}

func addAnimal(c *gin.Context) {
	var newAnimal Animal

	// Call BindJSON to bind the received JSON to newAnimal.
	if err := c.BindJSON(&newAnimal); err != nil {
		return
	}

	// Add the new animal to the slice.
	animals = append(animals, newAnimal)
	c.IndentedJSON(http.StatusCreated, newAnimal)
}

func updateAnimal(c *gin.Context) {
	id := c.Param("id")
	var updatedAnimal Animal

	// Call BindJSON to bind the received JSON to updatedAnimal.
	if err := c.BindJSON(&updatedAnimal); err != nil {
		return
	}

	for i, animal := range animals {
		if animal.ID == id {
			animals[i] = updatedAnimal
			c.IndentedJSON(http.StatusOK, updatedAnimal)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "animal not found"})
}

func deleteAnimal(c *gin.Context) {
	id := c.Param("id")
	for i, animal := range animals {
		if animal.ID == id {
			animals = append(animals[:i], animals[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "animal deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "animal not found"})
}

// Book represents data about a book.
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "1984", Author: "George Orwell"},
	{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee"},
}

// Handlers

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBook(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func addBook(c *gin.Context) {
	var newBook Book

	// Call BindJSON to bind the received JSON to newBook.
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add the new book to the slice.
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook Book

	// Call BindJSON to bind the received JSON to updatedBook.
	if err := c.BindJSON(&updatedBook); err != nil {
		return
	}

	for i, book := range books {
		if book.ID == id {
			books[i] = updatedBook
			c.IndentedJSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func main() {
	router := gin.Default()

	// Serve static files
	router.Static("/static", "./static")

	// Redirect from root path to /static/index.html
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/static/index.html")
	})

	// API routes
	api := router.Group("/api")
	{
		// Book routes
		api.GET("/books", getBooks)
		api.GET("/books/:id", getBook)
		api.POST("/books", addBook)
		api.PUT("/books/:id", updateBook)
		api.DELETE("/books/:id", deleteBook)

		// Animal routes
		api.GET("/animals", getAnimals)
		api.GET("/animals/:id", getAnimal)
		api.POST("/animals", addAnimal)
		api.PUT("/animals/:id", updateAnimal)
		api.DELETE("/animals/:id", deleteAnimal)
	}

	router.Run("localhost:8080")
}
