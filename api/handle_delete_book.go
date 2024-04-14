package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (cfg *apiConfig) deleteBook(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, err)
  }

	err = cfg.database.Delete(id)
  if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, err)
  }

	c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted"})
}
