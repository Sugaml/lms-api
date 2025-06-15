package domain

import "time"

type Book struct {
	BaseModel
	Title           string `gorm:"not null"`
	Author          string `gorm:"not null"`
	ISBN            string `gorm:"unique;not null"`
	Category        string `gorm:"not null"`
	Program         string `gorm:"not null"`
	TotalCopies     int    `gorm:"not null"`
	AvailableCopies int    `gorm:"not null"`
	Description     string
	CoverImage      string `gorm:"column:cover_image"`
}

type BookRequest struct {
	UserID       int        `json:"user_id"`
	BookID       int        `json:"book_id"`
	Status       string     `json:"status"` // 'pending' | 'approved' | 'rejected'
	RequestDate  time.Time  `json:"request_date"`
	ApprovedDate *time.Time `json:"approved_date"`
	ApprovedBy   *int       `json:"approved_by"`
}

type BookListRequest struct {
	ListRequest
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

func (r *BookUpdateRequest) NewUpdate() Map {
	return nil
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
