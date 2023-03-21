package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rmarasigan/go-bookstore/pkg/models"
	"github.com/rmarasigan/go-bookstore/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	// Converting it to JSON
	response, _ := json.Marshal(newBooks)

	// Return as JSON
	w.Header().Set("Content-Type", "application/json")
	// Returning 200 status
	w.WriteHeader(http.StatusOK)
	// Sending response to frontend
	w.Write(response)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	// Accessing parameters
	params := mux.Vars(r)
	bookId := params["bookId"]

	// Convering string to int
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	// Getting book details by id
	bookDetails, _ := models.GetBookById(id)

	// Sending json response
	response, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	// Parsing data to JSON
	utils.ParseBody(r, CreateBook)

	// Creating or inserting a new record to the database
	book := CreateBook.CreateBook()

	// Sending json response
	response, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]

	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	// Deleting book using id
	book := models.DeleteBook(id)

	// Sending json response
	response, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)

	params := mux.Vars(r)
	bookId := params["bookId"]

	// Convert string to int
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	// Find book using ID
	bookDetails, db := models.GetBookById(id)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)

	// Sending json response
	response, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
