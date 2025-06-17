package port

import (
	"github.com/sugaml/lms-api/internal/core/domain"
)

// type BorrowRepository interface is an interface for interacting with type Announcement-related data
type BorrowRepository interface {
	CreateBorrow(data *domain.BorrowedBook) (*domain.BorrowedBook, error)
	ListBorrow(req *domain.ListBorrowedBookRequest) ([]*domain.BorrowedBook, int64, error)
	GetBorrow(id string) (*domain.BorrowedBook, error)
	GetAvailableCopies(bookID string) (int, error)
	GetBookBorrowByUserID(user_id string) ([]*domain.BorrowedBook, error)
	IsBookBorrowByUserID(user_id string, book_id string) bool
	UpdateBorrow(id string, req domain.Map) (*domain.BorrowedBook, error)
	DeleteBorrow(id string) error
}

// type BorrowService interface is an interface for interacting with type Announcement-related data
type BorrowService interface {
	CreateBorrow(data *domain.BorrowedBookRequest) (*domain.BorrowedBookResponse, error)
	ListBorrow(req *domain.ListBorrowedBookRequest) ([]*domain.BorrowedBookResponse, int64, error)
	GetStudentsBorrowBook(id string) ([]*domain.BorrowedBookResponse, error)
	GetBorrow(id string) (*domain.BorrowedBookResponse, error)
	UpdateBorrow(id string, req *domain.UpdateBorrowedBookRequest) (*domain.BorrowedBookResponse, error)
	DeleteBorrow(id string) (*domain.BorrowedBookResponse, error)
}
