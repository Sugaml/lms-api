package domain

import "time"

type Fine struct {
	BaseModel
	UserID         int        `gorm:"not null"`
	BorrowedBookID int        `gorm:"column:borrowed_book_id;not null"`
	Amount         int        `gorm:"not null"` // in paisa
	Reason         string     `gorm:"not null"`
	Status         string     `gorm:"not null"` // 'pending' | 'paid'
	PaidAt         *time.Time `gorm:"column:paid_at"`
}

type FineRequest struct {
	ID             string     `json:"id"`
	CreatedAt      time.Time  `json:"created_at"`
	UserID         int        `json:"user_id"`
	BorrowedBookID int        `json:"borrowed_book_id"`
	Amount         int        `json:"amount"` // in paisa
	Reason         string     `json:"reason"`
	Status         string     `json:"status"` // 'pending' | 'paid'
	PaidAt         *time.Time `json:"paid_at"`
}

type UpdateFineRequest struct {
	UserID         int        `json:"user_id"`
	BorrowedBookID int        `json:"borrowed_book_id"`
	Amount         int        `json:"amount"` // in paisa
	Reason         string     `json:"reason"`
	Status         string     `json:"status"` // 'pending' | 'paid'
	PaidAt         *time.Time `json:"paid_at"`
}

type ListFineRequest struct {
	ListRequest
	UserID         int        `form:"user_id"`
	BorrowedBookID int        `form:"borrowed_book_id"`
	Amount         int        `form:"amount"` // in paisa
	Reason         string     `form:"reason"`
	Status         string     `form:"status"` // 'pending' | 'paid'
	PaidAt         *time.Time `form:"paid_at"`
}

type FineResponse struct {
	ID             string     `json:"id"`
	CreatedAt      time.Time  `json:"created_at"`
	UserID         int        `json:"user_id"`
	BorrowedBookID int        `json:"borrowed_book_id"`
	Amount         int        `json:"amount"` // in paisa
	Reason         string     `json:"reason"`
	Status         string     `json:"status"` // 'pending' | 'paid'
	PaidAt         *time.Time `json:"paid_at"`
}

func (u *FineRequest) Validate() error {
	return nil
}

func (r *UpdateFineRequest) NewUpdate() Map {
	return nil
}
