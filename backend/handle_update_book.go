package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/RealMotz/book-journal/internal/database"
	"github.com/gin-gonic/gin"
)

func (cfg *apiConfig) updateBook(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println("error parsing id")
		return
	}

	var book database.Book
	if err := c.BindJSON(&book); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	err = cfg.database.Update(id, book)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, struct{}{})
}
