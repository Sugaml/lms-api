package repository

import (
	"errors"

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
	err := f.Count(&count).
		Order(req.SortColumn + " " + req.SortDirection).
		Limit(req.Size).
		Offset(req.Size * (req.Page - 1)).
		Find(&datas).Error
	if err != nil {
		return nil, count, err
	}
	return datas, count, nil
}

func (r *Repository) GetBorrow(id string) (*domain.BorrowedBook, error) {
	var data domain.BorrowedBook
	if err := r.db.Model(&domain.BorrowedBook{}).
		Preload("BorrowedBook").
		Take(&data, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *Repository) UpdateBorrow(id string, req domain.Map) (*domain.BorrowedBook, error) {
	if id == "" {
		return nil, errors.New("required BorrowedBook id")
	}
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
