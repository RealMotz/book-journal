package main

import (
	"database/sql"
	"log"

	"github.com/RealMotz/book-journal/internal/database"
	"github.com/gin-gonic/gin"
)

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

type apiConfig struct {
	database *database.SqliteDB
}

func main() {
	sqliteDB := sql.DB{}
	db := database.SqliteDB{
		DB: &sqliteDB,
	}

	db.InitializeDB()
	config := apiConfig{
		database: &db,
	}

	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/books", config.getBooks)
	router.GET("/books/:id", config.getBook)
	router.POST("/books", config.createBook)
	router.PUT("/books/:id", config.updateBook)
	router.DELETE("/books/:id", config.deleteBook)
  err := router.Run()
  if err != nil {
    log.Fatal(err)
  }
}
