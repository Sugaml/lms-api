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

type BookListRequest struct {
	ListRequest
}

type BookAllUpdateRequest struct {
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
