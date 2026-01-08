package repository

import(
	"github.com/nassim-touissi/go-book-api/internal/model"
	"errors"
)

// Inexistant book handling
var ErrBookNotFound = errors.New("Book not found !")

// BookRepository defines CRUD operations for books
type BookRepository interface {
	Create(book *model.Book) error
	GetByID(id string) (*model.Book, error)
	List(limit, offset int) ([]*model.Book, error)
	Update(book *model.Book) error
	Delete(id string) error
}