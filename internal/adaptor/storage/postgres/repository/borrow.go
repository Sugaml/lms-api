package repository

import (
	"time"

	"github.com/sugaml/lms-api/internal/core/domain"
)

func (r *Repository) CreateBorrow(data *domain.BorrowedBook) (*domain.BorrowedBook, error) {
	if err := r.db.Model(&domain.BorrowedBook{}).Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) ListBorrow(req *domain.ListBorrowedBookRequest) ([]*domain.BorrowedBook, int64, error) {
	var datas []*domain.BorrowedBook
	var count int64
	f := r.db.Model(&domain.BorrowedBook{})
	if req.Query != "" {
		req.SortColumn = "score desc, " + req.SortColumn
	}
	if req.UserID != "" {
		f = f.Where("user_id = ?", req.UserID)
	}
	if req.BookID != "" {
		f = f.Where("book_id = ?", req.BookID)
	}
	if req.Status != "" {
		f = f.Where("status = ?", req.Status)
	}
	if req.BorrowedDate != (time.Time{}) {
		f = f.Where("borrowed_date = ?", req.BorrowedDate)
	}
	if req.DueDate != (time.Time{}) {
		f = f.Where("due_date = ?", req.DueDate)
	}
	err := f.Count(&count).
		Order(req.SortColumn + " " + req.SortDirection).
		Limit(req.Size).
		Offset(req.Size * (req.Page - 1)).
		Preload("Student").
		Preload("Book").
		Find(&datas).Error
	if err != nil {
		return nil, count, err
	}
	return datas, count, nil
}

func (r *Repository) GetBorrow(id string) (*domain.BorrowedBook, error) {
	var data domain.BorrowedBook
	if err := r.db.Model(&domain.BorrowedBook{}).
		Preload("Book").
		Preload("Student").
		Take(&data, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *Repository) GetBookBorrowByUserID(user_id string) ([]*domain.BorrowedBook, error) {
	var data []*domain.BorrowedBook
	if err := r.db.Model(&domain.BorrowedBook{}).
		Preload("Book").
		Preload("Student").
		Where("user_id = ?", user_id).Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) IsBookBorrowByUserID(userID string, bookID string) bool {
	var count int64
	err := r.db.Model(&domain.BorrowedBook{}).
		Where("user_id = ? AND book_id = ? AND returned_date IS NULL", userID, bookID).
		Count(&count).Error

	if err != nil {
		return false
	}
	return count > 0
}

func (r *Repository) UpdateBorrow(id string, req domain.Map) (*domain.BorrowedBook, error) {
	data := &domain.BorrowedBook{}
	err := r.db.Model(&domain.BorrowedBook{}).Where("id = ?", id).Updates(req.ToMap()).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) DeleteBorrow(id string) error {
	return r.db.Model(&domain.BorrowedBook{}).Where("id = ?", id).Delete(&domain.BorrowedBook{}).Error
}
