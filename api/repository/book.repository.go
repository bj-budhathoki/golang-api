package repository

import (
	"github.com/bj-budhathoki/golang-api/model"
	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(book model.Book) model.Book
	UpdateBook(book model.Book) model.Book
	DeleteBook(book model.Book)
	AllBook() []model.Book
	FindBookById(bookID uint64) model.Book
}
type bookConnection struct {
	connection *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookConnection{
		connection: db,
	}
}

func (db *bookConnection) CreateBook(book model.Book) model.Book {
	db.connection.Save(&book)
	db.connection.Preload("User").Find(&book)
	return book
}
func (db *bookConnection) UpdateBook(book model.Book) model.Book {
	db.connection.Save(&book)
	db.connection.Preload("User").Find(&book)
	return book
}
func (db *bookConnection) AllBook() []model.Book {
	var books []model.Book
	db.connection.Preload("User").Find(&books)
	return books
}

func (db *bookConnection) DeleteBook(book model.Book) {
	db.connection.Delete(&book)
}

func (db *bookConnection) FindBookById(bookID uint64) model.Book {
	var book model.Book
	db.connection.Preload("User").Find(&book, bookID)
	return book
}
