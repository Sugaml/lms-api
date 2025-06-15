package port

import "github.com/sugaml/lms-api/internal/core/domain"

type ReportRepository interface {
	GetBorrowedBookStats() (*domain.BorrowedBookStats, error)
}

type ReportService interface {
	GetBorrowedBookStats() (*domain.BorrowedBookStats, error)
}
