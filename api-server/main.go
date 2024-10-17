package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Book model
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "Book One", Author: "John Doe"},
	{ID: "2", Title: "Book Two", Author: "Steve Smith"},
}

func main() {
	r := gin.Default()

	// Create a group for API endpoints
	api := r.Group("/api/v1")
	{
		// GET all books
		api.GET("/books", getAllBooks)

		// GET a book by ID
		api.GET("/books/:id", getBookByID)

		// POST a new book
		api.POST("/books", createBook)

		// PUT to update a book
		api.PUT("/books/:id", updateBook)

		// DELETE a book
		api.DELETE("/books/:id", deleteBook)
	}

	r.Run(":8000")
}

// GET all books
func getAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

// GET a book by ID
func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

// POST a new book
func createBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

// PUT to update a book
func updateBook(c *gin.Context) {
	id := c.Param("id")

	for index, book := range books {
		if book.ID == id {
			var updatedBook Book

			if err := c.BindJSON(&updatedBook); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
				return
			}

			books[index] = updatedBook
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

// DELETE a book
func deleteBook(c *gin.Context) {
	id := c.Param("id")

	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}
