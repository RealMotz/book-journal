package main

import (
	"net/http"

	"github.com/RealMotz/book-journal/internal/database"
	"github.com/gin-gonic/gin"
)

func (cfg *apiConfig) updateBook(c *gin.Context) {
	var book database.Book
	if err := c.BindJSON(&book); err != nil {
    c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

  err := cfg.database.Update(book)
  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, err)
		return
  }

	c.IndentedJSON(http.StatusOK, struct{}{})
}
