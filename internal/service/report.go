package service

import "github.com/sugaml/lms-api/internal/core/domain"

func (s *Service) GetBorrowedBookStats() (*domain.BorrowedBookStats, error) {
	result, err := s.repo.GetBorrowedBookStats()
	if err != nil {
		return nil, err
	}
	return result, nil
}
