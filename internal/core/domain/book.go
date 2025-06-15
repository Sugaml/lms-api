package domain

import (
	"errors"
	"time"
)

type Book struct {
	BaseModel
	Title       string `gorm:"not null" json:"title"`
	Author      string `gorm:"not null" json:"author"`
	ISBN        string `gorm:"unique;not null" json:"isbn"`
	Category    string `gorm:"not null" json:"category"`
	CategoryID  string `gorm:"not null" json:"category_id"`
	Program     string `gorm:"not null" json:"program"`
	TotalCopies int    `gorm:"not null" json:"total_copies"`
	Description string `gorm:"not null" json:"description"`
	CoverImage  string `gorm:"column:cover_image" json:"cover_image"`
	IsActive    bool   `gorm:"column:is_active;default:false" json:"is_active"`
}

type BookRequest struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Category    string `json:"category"`
	Program     string `json:"program"`
	TotalCopies int    `json:"total_copies"`
	Description string `json:"description"`
	CoverImage  string `json:"cover_image"`
}

type BookListRequest struct {
	ListRequest
	Title       string `form:"title"`
	Author      string `form:"author"`
	ISBN        string `form:"isbn"`
	Category    string `form:"category"`
	Program     string `form:"program"`
	TotalCopies int    `form:"total_copies"`
	Description string `form:"description"`
	CoverImage  string `form:"cover_image"`
}

type BookUpdateRequest struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Category    string `json:"category"`
	Program     string `json:"program"`
	TotalCopies int    `json:"total_copies"`
	Description string `json:"description"`
	CoverImage  string `json:"cover_image"`
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
	ID              string    `json:"id"`
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
	if r.Title == "" {
		return errors.New("title is required")
	}
	if r.Author == "" {
		return errors.New("author is required")
	}
	if r.ISBN == "" {
		return errors.New("isbn is required")
	}
	return nil
}

func (r *BookUpdateRequest) NewUpdate() Map {
	return nil
}
