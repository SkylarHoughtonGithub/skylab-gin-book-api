package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Book Handlers

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func getBooks(c *gin.Context) {
	var books []Book
	if err := readJSONFile(booksFile, &books); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to read books"})
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

func getBook(c *gin.Context) {
	indexStr := c.Param("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid index"})
		return
	}
	var books []Book
	if err := readJSONFile(booksFile, &books); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to read books"})
		return
	}

	if index >= 0 && index < len(books) {
		c.IndentedJSON(http.StatusOK, books[index])
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
	}
}

func addBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	var books []Book
	if err := readJSONFile(booksFile, &books); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to read books"})
		return
	}
	books = append(books, newBook)
	if err := writeJSONFile(booksFile, books); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to save book"})
		return
	}
	c.IndentedJSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
	indexStr := c.Param("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid index"})
		return
	}
	var updatedBook Book
	if err := c.BindJSON(&updatedBook); err != nil {
		return
	}
	var books []Book
	if err := readJSONFile(booksFile, &books); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to read books"})
		return
	}

	if index >= 0 && index < len(books) {
		books[index] = updatedBook
		if err := writeJSONFile(booksFile, books); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to save book"})
			return
		}
		c.IndentedJSON(http.StatusOK, updatedBook)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
	}
}

func deleteBook(c *gin.Context) {
	indexStr := c.Param("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid index"})
		return
	}
	var books []Book
	if err := readJSONFile(booksFile, &books); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to read books"})
		return
	}

	if index >= 0 && index < len(books) {
		books = append(books[:index], books[index+1:]...)
		if err := writeJSONFile(booksFile, books); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to save books"})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Book deleted"})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
	}
}
