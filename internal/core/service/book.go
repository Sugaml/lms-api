package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/sugaml/lms-api/internal/core/domain"
)

// CreateBook creates a new Book
func (s *Service) CreateBook(ctx context.Context, req *domain.BookRequest) (*domain.BookResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	data := domain.Convert[domain.BookRequest, domain.Book](req)
	result, err := s.repo.CreateBook(data)
	if err != nil {
		return nil, err
	}
	userID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, _ = s.repo.CreateNotification(&domain.Notification{
		Title:       fmt.Sprintf("Created new %s Book.", result.Title),
		Description: "create",
		UserID:      userID,
		Type:        "book",
		Action:      "create",
		Module:      "book",
		IsActive:    true,
	})
	_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    fmt.Sprintf("Created new Book %s.", result.Title),
		Action:   "create",
		Data:     string(domain.ConvertToJson(result)),
		IsActive: true,
	})
	return domain.Convert[domain.Book, domain.BookResponse](result), nil
}

// ListBooks retrieves a list of Books
func (s *Service) ListBook(ctx context.Context, req *domain.BookListRequest) ([]*domain.BookResponse, int64, error) {
	var datas = []*domain.BookResponse{}
	results, count, err := s.repo.ListBook(req)
	if err != nil {
		return nil, count, err
	}
	for _, result := range results {
		data := domain.Convert[domain.Book, domain.BookResponse](result)
		data.AvailableCopies, _ = s.repo.GetAvailableCopies(data.ID)
		datas = append(datas, data)
	}
	return datas, count, nil
}

func (s *Service) GetBook(ctx context.Context, id string) (*domain.BookResponse, error) {
	result, err := s.repo.GetBook(id)
	if err != nil {
		return nil, err
	}
	data := domain.Convert[domain.Book, domain.BookResponse](result)
	return data, nil
}

func (s *Service) UpdateBook(ctx context.Context, id string, req *domain.BookUpdateRequest) (*domain.BookResponse, error) {
	if id == "" {
		return nil, errors.New("required Book id")
	}
	_, err := s.repo.GetBook(id)
	if err != nil {
		return nil, err
	}

	// update
	mp := req.NewUpdate()
	result, err := s.repo.UpdateBook(id, mp)
	if err != nil {
		return nil, err
	}
	_, _ = s.repo.CreateNotification(&domain.Notification{
		Title:       fmt.Sprintf("Updated %s Book details.", result.Title),
		Description: "update",
		Type:        "book",
		Action:      "update",
		Module:      "book",
		IsActive:    true,
	})
	_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    fmt.Sprintf("Updated %s Book details.", result.Title),
		Action:   "update",
		Data:     fmt.Sprint(req),
		IsActive: true,
	})
	data := domain.Convert[domain.Book, domain.BookResponse](result)
	return data, nil
}

func (s *Service) DeleteBook(ctx context.Context, id string) (*domain.BookResponse, error) {
	result, err := s.repo.GetBook(id)
	if err != nil {
		return nil, err
	}
	CountBorrwedCopiesBookID, err := s.repo.CountBorrwedCopiesBookID(id)
	if err != nil {
		return nil, err
	}
	logrus.Info("CountBorrwedCopiesBookID :: ", CountBorrwedCopiesBookID)
	if CountBorrwedCopiesBookID > 0 {
		return nil, fmt.Errorf("book has %d copies borrowed cannot delete it", CountBorrwedCopiesBookID)
	}
	err = s.repo.DeleteBook(id)
	if err != nil {
		return nil, err
	}
	_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    fmt.Sprintf("Deleted %s parking area.", result.Title),
		Action:   "delete",
		Data:     fmt.Sprint(result),
		IsActive: true,
	})
	return domain.Convert[domain.Book, domain.BookResponse](result), nil
}
