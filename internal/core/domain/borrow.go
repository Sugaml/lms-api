package domain

import (
	"errors"
	"time"
)

type BorrowedBook struct {
	BaseModel
	UserID       string     `gorm:"not null" json:"user_id"`
	BookID       string     `gorm:"not null" json:"book_id"`
	LibrarianID  string     `json:"librarian_id"`
	BorrowedDate time.Time  `gorm:"column:borrowed_date;autoCreateTime" json:"borrowed_date"`
	DueDate      time.Time  `gorm:"not null" json:"due_date"`
	ReturnedDate *time.Time `gorm:"column:returned_date" json:"returned_date"`
	RenewalCount int        `gorm:"default:0" json:"renewal_count"`
	Status       string     `gorm:"not null" json:"status"` // 'borrowed' | 'returned' | 'overdue'
	Remarks      string     `json:"remarks"`
	IsActive     bool       `gorm:"column:is_active;default:false" json:"is_active"`
	Student      *User      `gorm:"foreignkey:ID;references:UserID" json:"student"`
	Librarian    *User      `gorm:"foreignkey:ID;references:LibrarianID" json:"librarian"`
	Book         *Book      `gorm:"foreignkey:ID;references:BookID" json:"book"`
}

type BorrowedBookRequest struct {
	UserID       string     `json:"user_id"`
	BookID       string     `json:"book_id"`
	DueDate      time.Time  `json:"due_date"`
	ReturnedDate *time.Time `json:"returned_date"`
	RenewalCount int        `json:"renewal_count"`
}

type UpdateBorrowedBookRequest struct {
	UserID       string     `json:"user_id"`
	BookID       string     `json:"book_id"`
	BorrowedDate time.Time  `json:"borrowed_date"`
	DueDate      time.Time  `json:"due_date"`
	LibrarianID  string     `json:"librarian_id"`
	ReturnedDate *time.Time `json:"returned_date"`
	RenewalCount int        `json:"renewal_count"`
	Remarks      string     `json:"remarks"`
	Status       string     `json:"status"` // 'borrowed' | 'returned' | 'overdue'
}

type ListBorrowedBookRequest struct {
	ListRequest
	UserID       string     `form:"user_id"`
	BookID       string     `form:"book_id"`
	LibrarianID  string     `json:"librarian_id"`
	BorrowedDate time.Time  `form:"borrowed_date"`
	DueDate      time.Time  `form:"due_date"`
	ReturnedDate *time.Time `form:"returned_date"`
	RenewalCount int        `form:"renewal_count"`
	Status       string     `form:"status"` // 'pending' | ''borrowed' | 'returned' | 'overdue'
}

type BorrowedBookResponse struct {
	ID           string        `json:"id"`
	CreatedAt    time.Time     `json:"created_at"`
	UserID       string        `json:"user_id"`
	BookID       string        `json:"book_id"`
	LibrarianID  string        `json:"librarian_id"`
	BorrowedDate time.Time     `json:"borrowed_date"`
	DueDate      time.Time     `json:"due_date"`
	ReturnedDate *time.Time    `json:"returned_date"`
	RenewalCount int           `json:"renewal_count"`
	Status       string        `json:"status"` // 'borrowed' | 'returned' | 'overdue'
	Student      *UserResponse `json:"student"`
	Librarian    *UserResponse `json:"librarian"`
	Book         *BookResponse `json:"book"`
	Remarks      string        `json:"remarks"`
}

func (r BorrowedBookRequest) Validate() error {
	if r.UserID == "" {
		return errors.New("user id is required")
	}
	if r.BookID == "" {
		return errors.New("book id is required")
	}
	if r.DueDate.IsZero() {
		return errors.New("due date is required")
	}
	return nil
}

func (r *UpdateBorrowedBookRequest) NewUpdate() Map {
	mp := map[string]interface{}{}
	if r.UserID != "" {
		mp["user_id"] = r.UserID
	}
	if r.BookID != "" {
		mp["book_id"] = r.BookID
	}
	if !r.DueDate.IsZero() {
		mp["due_date"] = r.DueDate
	}
	if r.LibrarianID != "" {
		mp["librarian_id"] = r.LibrarianID
	}
	if r.Remarks != "" {
		mp["remarks"] = r.Remarks
	}
	if r.ReturnedDate != nil {
		mp["returned_date"] = r.ReturnedDate
	}
	if r.RenewalCount != 0 {
		mp["renewal_count"] = r.RenewalCount
	}
	if r.Status != "" {
		if r.Status == "borrowed" {
			mp["borrowed_date"] = time.Now()
		}
		if r.Status == "returned" {
			mp["returned_date"] = time.Now()
		}
		mp["status"] = r.Status
	}
	return mp
}
