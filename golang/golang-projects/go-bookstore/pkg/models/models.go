package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rmarasigan/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// Initialize database connection
func init() {
	config.Connect()
	db = config.GetDB()

	// auto migration for given models, will only add
	// missing fields, won't delete/change current data.
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	// Check if PK is blank
	db.NewRecord(b)
	// Insert new record to the database
	db.Create(&b)
	return b
}

// Returning slice of books or list of books
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	// Running a where command using the ID to find the book
	db := db.Where("ID = ?", Id).Find(&getBook)

	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(book)
	return book
}
