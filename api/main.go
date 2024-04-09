package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Notes       string `json:"notes"`
}

type bookResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Notes       string `json:"notes"`
}

var library = []book{
	{Id: "1", Name: "Awesome book 1", Description: "This book is cool 1"},
	{Id: "2", Name: "Awesome book 2", Description: "This book is cool 2"},
	{Id: "3", Name: "Awesome book 3", Description: "This book is cool 3"},
	{Id: "4", Name: "Awesome book 4", Description: "This book is cool 4"},
	{Id: "5", Name: "Awesome book 5", Description: "This book is cool 5"},
	{Id: "6", Name: "Awesome book 6", Description: "This book is cool 6"},
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
		Name:        newBook.Name,
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

func main() {
	fmt.Println("Books API here")
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBook)
	router.POST("/books", createBook)
  router.DELETE("/books/:id", deleteBook)
	router.Run()
}
