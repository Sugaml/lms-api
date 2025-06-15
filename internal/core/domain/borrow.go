package domain

import "time"

type BorrowedBook struct {
	BaseModel
	UserID       int        `gorm:"not null" json:"user_id"`
	BookID       int        `gorm:"not null" json:"book_id"`
	BorrowedDate time.Time  `gorm:"column:borrowed_date;autoCreateTime" json:"borrowed_date"`
	DueDate      time.Time  `gorm:"not null" json:"due_date"`
	ReturnedDate *time.Time `gorm:"column:returned_date" json:"returned_date"`
	RenewalCount int        `gorm:"default:0" json:"renewal_count"`
	Status       string     `gorm:"not null" json:"status"` // 'borrowed' | 'returned' | 'overdue'
	IsActive     bool       `gorm:"column:is_active;default:false" json:"is_active"`
}

type BorrowedBookRequest struct {
	UserID       string     `json:"user_id"`
	BookID       string     `json:"book_id"`
	BorrowedDate time.Time  `json:"borrowed_date"`
	DueDate      time.Time  `json:"due_date"`
	ReturnedDate *time.Time `json:"returned_date"`
	RenewalCount int        `json:"renewal_count"`
	Status       string     `json:"status"` // 'borrowed' | 'returned' | 'overdue'
}

type UpdateBorrowedBookRequest struct {
	UserID       string     `json:"user_id"`
	BookID       string     `json:"book_id"`
	BorrowedDate time.Time  `json:"borrowed_date"`
	DueDate      time.Time  `json:"due_date"`
	ReturnedDate *time.Time `json:"returned_date"`
	RenewalCount int        `json:"renewal_count"`
	Status       string     `json:"status"` // 'borrowed' | 'returned' | 'overdue'
}

type ListBorrowedBookRequest struct {
	ListRequest
	UserID       string     `form:"user_id"`
	BookID       string     `form:"book_id"`
	BorrowedDate time.Time  `form:"borrowed_date"`
	DueDate      time.Time  `form:"due_date"`
	ReturnedDate *time.Time `form:"returned_date"`
	RenewalCount int        `form:"renewal_count"`
	Status       string     `form:"status"` // 'borrowed' | 'returned' | 'overdue'
}

type BorrowedBookResponse struct {
	ID           string     `json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UserID       string     `json:"user_id"`
	BookID       string     `json:"book_id"`
	BorrowedDate time.Time  `json:"borrowed_date"`
	DueDate      time.Time  `json:"due_date"`
	ReturnedDate *time.Time `json:"returned_date"`
	RenewalCount int        `json:"renewal_count"`
	Status       string     `json:"status"` // 'borrowed' | 'returned' | 'overdue'
}

func (r BorrowedBookRequest) Validate() error {
	return nil
}

func (r *UpdateBorrowedBookRequest) NewUpdate() Map {
	return nil
}
