package service

import (
	"errors"

	"github.com/sugaml/lms-api/internal/core/domain"
)

// CreateBorrowBook creates a new BorrowedBook
func (s *Service) CreateBorrow(req *domain.BorrowedBookRequest) (*domain.BorrowedBookResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	data := domain.Convert[domain.BorrowedBookRequest, domain.BorrowedBook](req)
	result, err := s.repo.CreateBorrow(data)
	if err != nil {
		return nil, err
	}
	return domain.Convert[domain.BorrowedBook, domain.BorrowedBookResponse](result), nil
}

// ListBorrowedBooks retrieves a list of BorrowedBooks
func (s *Service) ListBorrow(req *domain.ListBorrowedBookRequest) ([]*domain.BorrowedBookResponse, int64, error) {
	var datas = []*domain.BorrowedBookResponse{}
	results, count, err := s.repo.ListBorrow(req)
	if err != nil {
		return nil, count, err
	}
	for _, result := range results {
		data := domain.Convert[domain.BorrowedBook, domain.BorrowedBookResponse](result)
		datas = append(datas, data)
	}
	return datas, count, nil
}

func (s *Service) GetBorrow(id string) (*domain.BorrowedBookResponse, error) {
	result, err := s.repo.GetBorrow(id)
	if err != nil {
		return nil, err
	}
	data := domain.Convert[domain.BorrowedBook, domain.BorrowedBookResponse](result)
	return data, nil
}

func (s *Service) UpdateBorrow(id string, req *domain.UpdateBorrowedBookRequest) (*domain.BorrowedBookResponse, error) {
	if id == "" {
		return nil, errors.New("required BorrowedBook id")
	}
	_, err := s.repo.GetBorrow(id)
	if err != nil {
		return nil, err
	}
	// update
	mp := req.NewUpdate()
	result, err := s.repo.UpdateBorrow(id, mp)
	if err != nil {
		return nil, err
	}
	data := domain.Convert[domain.BorrowedBook, domain.BorrowedBookResponse](result)
	return data, nil
}

func (s *Service) DeleteBorrow(id string) (*domain.BorrowedBookResponse, error) {
	result, err := s.repo.GetBorrow(id)
	if err != nil {
		return nil, err
	}
	err = s.repo.DeleteBorrow(id)
	if err != nil {
		return nil, err
	}
	return domain.Convert[domain.BorrowedBook, domain.BorrowedBookResponse](result), nil
}
