// internal/handlers/book_handlers.go

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"skylab-gin-book-api/internal/database"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	DB *database.DB
}

func NewBookHandler(db *database.DB) *BookHandler {
	return &BookHandler{DB: db}
}

func (h *BookHandler) GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	
	// Try to get the book from cache
	cacheKey := fmt.Sprintf("book:%d", id)
	cachedBook, err := h.DB.GetCache(cacheKey)
	if err == nil {
		var book database.Book
		json.Unmarshal([]byte(cachedBook), &book)
		c.JSON(http.StatusOK, book)
		return
	}

	// If not in cache, get from database
	book, err := h.DB.GetBook(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Cache the book
	bookJSON, _ := json.Marshal(book)
	h.DB.SetCache(cacheKey, bookJSON, time.Hour)

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var book database.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.CreateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	// Cache the new book
	cacheKey := fmt.Sprintf("book:%d", book.ID)
	bookJSON, _ := json.Marshal(book)
	h.DB.SetCache(cacheKey, bookJSON, time.Hour)

	c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	var book database.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.UpdateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	// Update the book in cache
	cacheKey := fmt.Sprintf("book:%d", book.ID)
	bookJSON, _ := json.Marshal(book)
	h.DB.SetCache(cacheKey, bookJSON, time.Hour)

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.DB.DeleteBook(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	// Remove the book from cache
	cacheKey := fmt.Sprintf("book:%d", id)
	h.DB.DeleteCache(cacheKey)

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

func (h *BookHandler) ListBooks(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	// Try to get the book list from cache
	cacheKey := fmt.Sprintf("books:list:%d:%d", limit, offset)
	cachedBooks, err := h.DB.GetCache(cacheKey)
	if err == nil {
		var books []database.Book
		json.Unmarshal([]byte(cachedBooks), &books)
		c.JSON(http.StatusOK, books)
		return
	}

	// If not in cache, get from database
	books, err := h.DB.ListBooks(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}

	// Cache the book list
	booksJSON, _ := json.Marshal(books)
	h.DB.SetCache(cacheKey, booksJSON, time.Hour)

	c.JSON(http.StatusOK, books)
}