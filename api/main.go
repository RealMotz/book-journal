package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Notes       string `json:"notes"`
}

type bookResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Notes       string `json:"notes"`
}

var library = []book{
	{Id: "1", Title: "Awesome book 1", Description: "This book is cool 1"},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, library)
}

func getBook(c *gin.Context) {
	id := c.Param("id")
	for _, book := range library {
		if book.Id == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		log.Fatal(err)
		return
	}

	newBook.Id = fmt.Sprint(len(library) + 1)

	library = append(library, newBook)
	c.IndentedJSON(http.StatusCreated, bookResponse{
		Title:       newBook.Title,
		Description: newBook.Description,
		Notes:       newBook.Notes,
	})
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	for i, book := range library {
		if book.Id == id {
			deleteItem(library, i)
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func deleteItem(slice []book, index int) []book {
	newSlice := make([]book, 0)
	newSlice = append(newSlice, slice[:index]...)
	return append(newSlice, slice[index+1:]...)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	fmt.Println("Books API here")
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBook)
	router.POST("/books", createBook)
	router.DELETE("/books/:id", deleteBook)
	router.Run()
}
