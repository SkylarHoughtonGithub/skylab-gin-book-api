package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Animal Handlers

type Animal struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

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
