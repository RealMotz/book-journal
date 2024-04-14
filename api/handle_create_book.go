package main

import (
	"net/http"

	"github.com/RealMotz/book-journal/internal/database"
	"github.com/gin-gonic/gin"
)

func (cfg *apiConfig) createBook(c *gin.Context) {
	var newBook database.Book
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error parsing body"})
		return
	}

	type bookResponse struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Memo        string `json:"memo"`
		IsReading   bool   `json:"isReading"`
	}

	err := cfg.database.Insert(newBook)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, bookResponse{
		Name:        newBook.Name,
		Description: newBook.Description,
		Memo:        newBook.Memo,
		IsReading:   newBook.IsReading,
	})
}
