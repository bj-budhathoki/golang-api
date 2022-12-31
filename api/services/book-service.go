package services

import (
	"fmt"
	"log"

	"github.com/bj-budhathoki/golang-api/api/repository"
	"github.com/bj-budhathoki/golang-api/dtos"
	"github.com/bj-budhathoki/golang-api/model"
	"github.com/mashingan/smapping"
)

type BookService interface {
	CreateBook(book dtos.BookCreateDTOS) model.Book
	UpdateBook(book dtos.BookUpdateDTOS) model.Book
	DeleteBook(book model.Book)
	AllBook() []model.Book
	FindBookById(bookID uint64) model.Book
	IsAllowedToEdit(userID string, bookID uint64) bool
}
type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepo,
	}
}

func (b *bookService) CreateBook(book dtos.BookCreateDTOS) model.Book {
	bookToCreate := model.Book{}
	err := smapping.FillStruct(&bookToCreate, smapping.MapFields(&book))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := b.bookRepository.CreateBook(bookToCreate)
	return res
}

func (b *bookService) UpdateBook(book dtos.BookUpdateDTOS) model.Book {
	bookToUpdate := model.Book{}
	err := smapping.FillStruct(&bookToUpdate, smapping.MapFields(book))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	updatedBook := b.bookRepository.UpdateBook(bookToUpdate)
	return updatedBook
}
func (s *bookService) AllBook() []model.Book {
	return s.bookRepository.AllBook()
}
func (s *bookService) DeleteBook(b model.Book) {
	s.bookRepository.DeleteBook(b)
}
func (s *bookService) FindBookById(bookID uint64) model.Book {
	return s.bookRepository.FindBookById(bookID)
}
func (service *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
	b := service.bookRepository.FindBookById(bookID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
