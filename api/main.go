package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
  Name string `json:"name"`
  Description string `json:"description"`
}

func main() {
	fmt.Println("Books API here")
	server := gin.Default()
	server.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"result": []book{
				{
          Name: "Awesome book 1",
					Description: "This is an awesome book",
        },
				{
          Name: "Awesome book 2",
					Description: "This is an awesome book 2",
        },
			},
		})
	})
	server.Run()
}
