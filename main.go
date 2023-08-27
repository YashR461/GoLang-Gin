package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
)
type book struct {
	ID string		`json:"id"`
	Title string	`json:"title"`
	Author string	`json:"author"`
	Quantity int	`json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context){
	id := c.Param("id")
	book , err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK , book)
}
func getBookById(id string) (*book , error) {
	for i:=0 ; i<len(books) ; i++ {
		if books[i].ID == id {
			return &books[i] , nil
		}
	}
	return nil , errors.New("book not found")
}

func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook) ; err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func issueBook(c *gin.Context) {
	id , ok := c.GetQuery("id")
 
	if ok != false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"mesage":"Missing id query parameter"})
		return
	}
	book , err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Book not found"})
		return
	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Book not available"})
		return
	}
	book.Quantity--
	c.IndentedJSON(http.StatusOK, book)

}

func main() {

	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books" , createBook)
	router.GET("/books/:id" , bookById)
	router.PATCH("/issue", issueBook)

	router.Run(":8080")

}