package port

import (
	"github.com/sugaml/lms-api/internal/core/domain"
)

// type BookRepository interface is an interface for interacting with type Announcement-related data
type BookRepository interface {
	CreateBook(data *domain.Book) (*domain.Book, error)
	ListBook(req *domain.BookListRequest) ([]*domain.Book, int64, error)
	GetBook(id string) (*domain.Book, error)
	UpdateBook(id string, req domain.Map) (*domain.Book, error)
	DeleteBook(id string) error
}

// type BookService interface is an interface for interacting with type Announcement-related data
type BookService interface {
	CreateBook(data *domain.BookRequest) (*domain.BookResponse, error)
	ListBook(req *domain.BookListRequest) ([]*domain.BookResponse, int64, error)
	GetBook(id string) (*domain.BookResponse, error)
	UpdateBook(id string, req *domain.BookUpdateRequest) (*domain.BookResponse, error)
	DeleteBook(id string) (*domain.BookResponse, error)
}
