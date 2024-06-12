package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

func getBooks(c *gin.Context) {
	var books []book
	DB.Find(&books)
	c.IndentedJSON(http.StatusOK, books)
}

func getBookById(c *gin.Context) {
	var book book
	id := c.Param("id")
	asdf := DB.Find(&book, "id = ?", id)
	if asdf.RowsAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func updateBook(c *gin.Context) {
	var jsonData book
	c.BindJSON(&jsonData)

	var book book
	DB.Find(&book, "id = ?", c.Param("id"))
	book.Quantity = jsonData.Quantity
	DB.Save(&book)
	c.IndentedJSON(http.StatusOK, &book)
}

func deleteBook(c *gin.Context) {
	var book book
	DB.Find(&book, "id = ?", c.Param("id"))
	DB.Delete(&book)
	c.IndentedJSON(http.StatusOK, nil)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		fmt.Println("FUKC", err)
		return
	}
	fmt.Println(newBook)

	DB.Create(&newBook)
	c.JSON(http.StatusCreated, newBook)
}

var DB *gorm.DB

func connectDB() {
	dbURL := "postgres://skhan:haha@localhost:5432/book_store"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&book{})
	DB = db
}

func main() {
	connectDB()

	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", createBook)
	router.PATCH("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)
	router.Run("localhost:8080")
}
