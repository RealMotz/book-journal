package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (cfg *apiConfig) getBook(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println("error fetching params?")
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	book, err := cfg.database.Select(id)
	if err != nil {
		fmt.Println(err)
		fmt.Println("error fetching book")
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func (cfg *apiConfig) getBooks(c *gin.Context) {
	reading := c.Query("reading")
	books, err := cfg.database.SelectAll(reading == "true")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, struct{}{})
		return
	}

	c.IndentedJSON(http.StatusOK, books)
}
