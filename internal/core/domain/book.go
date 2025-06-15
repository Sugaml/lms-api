package domain

import (
	"errors"
	"time"
)

type Book struct {
	BaseModel
	Title           string `gorm:"not null" json:"title"`
	Author          string `gorm:"not null" json:"author"`
	ISBN            string `gorm:"unique;not null" json:"isbn"`
	Category        string `gorm:"not null" json:"category"`
	CategoryID      string `gorm:"not null" json:"category_id"`
	Program         string `gorm:"not null" json:"program"`
	TotalCopies     int    `gorm:"not null" json:"total_copies"`
	AvailableCopies int    `gorm:"not null" json:"available_copies"`
	Description     string `gorm:"not null" json:"description"`
	CoverImage      string `gorm:"column:cover_image" json:"cover_image"`
	IsActive        bool   `gorm:"column:is_active;default:false" json:"is_active"`
}

type BookRequest struct {
	UserID       int        `json:"user_id"`
	BookID       int        `json:"book_id"`
	CategoryID   string     `json:"category_id"`
	Status       string     `json:"status"` // 'pending' | 'approved' | 'rejected'
	RequestDate  time.Time  `json:"request_date"`
	ApprovedDate *time.Time `json:"approved_date"`
	ApprovedBy   *int       `json:"approved_by"`
}

type BookListRequest struct {
	ListRequest
	Title           string `form:"title"`
	Author          string `form:"author"`
	ISBN            string `form:"isbn"`
	Category        string `form:"category"`
	Program         string `form:"program"`
	TotalCopies     int    `form:"total_copies"`
	AvailableCopies int    `form:"available_copies"`
	Description     string `form:"description"`
	CoverImage      string `form:"cover_image"`
}

type BookUpdateRequest struct {
	Title           string `json:"title"`
	Author          string `json:"author"`
	ISBN            string `json:"isbn"`
	Category        string `json:"category"`
	Program         string `json:"program"`
	TotalCopies     int    `json:"total_copies"`
	AvailableCopies int    `json:"available_copies"`
	Description     string `json:"description"`
	CoverImage      string `json:"cover_image"`
}

type BookAllUpdateRequest struct {
	Title           string `json:"title"`
	Author          string `json:"author"`
	ISBN            string `json:"isbn"`
	Category        string `json:"category"`
	Program         string `json:"program"`
	TotalCopies     int    `json:"total_copies"`
	AvailableCopies int    `json:"available_copies"`
	Description     string `json:"description"`
	CoverImage      string `json:"cover_image"`
}

type BookResponse struct {
	ID              int       `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	ISBN            string    `json:"isbn"`
	Category        string    `json:"category"`
	Program         string    `json:"program"`
	TotalCopies     int       `json:"total_copies"`
	AvailableCopies int       `json:"available_copies"`
	Description     string    `json:"description"`
	CoverImage      string    `json:"cover_image"`
}

func (r *BookRequest) Validate() error {
	if r.UserID == 0 {
		return errors.New("user id is required")
	}
	if r.BookID == 0 {
		return errors.New("book id is required")
	}
	if r.Status == "" {
		return errors.New("status is required")
	}
	if r.RequestDate.IsZero() {
		return errors.New("request date is required")
	}
	return nil
}

func (r *BookUpdateRequest) NewUpdate() Map {
	return nil
}
