package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	booksFile   = "books.json"
	animalsFile = "animals.json"
	mu          sync.Mutex
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Animal struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func readJSONFile(filename string, v interface{}) error {
	mu.Lock()
	defer mu.Unlock()
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist, return empty data
			return nil
		}
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

func writeJSONFile(filename string, v interface{}) error {
	mu.Lock()
	defer mu.Unlock()
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	return err
}

// Book Handlers
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

// Animal Handlers
func getAnimals(c *gin.Context) {
	var animals []Animal
	if err := readJSONFile(animalsFile, &animals); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to read animals"})
		return
	}
	c.IndentedJSON(http.StatusOK, animals)
}

func getAnimal(c *gin.Context) {
	indexStr := c.Param("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid index"})
		return
	}
	var animals []Animal
	if err := readJSONFile(animalsFile, &animals); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to read animals"})
		return
	}

	if index >= 0 && index < len(animals) {
		c.IndentedJSON(http.StatusOK, animals[index])
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Animal not found"})
	}
}

func addAnimal(c *gin.Context) {
	var newAnimal Animal
	if err := c.BindJSON(&newAnimal); err != nil {
		return
	}
	var animals []Animal
	if err := readJSONFile(animalsFile, &animals); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to read animals"})
		return
	}
	animals = append(animals, newAnimal)
	if err := writeJSONFile(animalsFile, animals); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to save animal"})
		return
	}
	c.IndentedJSON(http.StatusCreated, newAnimal)
}

func updateAnimal(c *gin.Context) {
	indexStr := c.Param("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid index"})
		return
	}
	var updatedAnimal Animal
	if err := c.BindJSON(&updatedAnimal); err != nil {
		return
	}
	var animals []Animal
	if err := readJSONFile(animalsFile, &animals); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to read animals"})
		return
	}

	if index >= 0 && index < len(animals) {
		animals[index] = updatedAnimal
		if err := writeJSONFile(animalsFile, animals); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to save animal"})
			return
		}
		c.IndentedJSON(http.StatusOK, updatedAnimal)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Animal not found"})
	}
}

func deleteAnimal(c *gin.Context) {
	indexStr := c.Param("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid index"})
		return
	}
	var animals []Animal
	if err := readJSONFile(animalsFile, &animals); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to read animals"})
		return
	}

	if index >= 0 && index < len(animals) {
		animals = append(animals[:index], animals[index+1:]...)
		if err := writeJSONFile(animalsFile, animals); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to save animals"})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Animal deleted"})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Animal not found"})
	}
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
		api.GET("/books/:index", getBook)
		api.POST("/books", addBook)
		api.PUT("/books/:index", updateBook)
		api.DELETE("/books/:index", deleteBook)

		// Animal routes
		api.GET("/animals", getAnimals)
		api.GET("/animals/:index", getAnimal)
		api.POST("/animals", addAnimal)
		api.PUT("/animals/:index", updateAnimal)
		api.DELETE("/animals/:index", deleteAnimal)
	}

	// Start server
	router.Run(":8080")
}
