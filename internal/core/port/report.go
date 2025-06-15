package port

import "github.com/sugaml/lms-api/internal/core/domain"

type ReportRepository interface {
	GetLibraryDashboardStats() (*domain.LibraryDashboardStats, error)
	GetBorrowedBookStats() (*domain.BorrowedBookStats, error)
}

type ReportService interface {
	GetLibraryDashboardStats() (*domain.LibraryDashboardStats, error)
	GetBorrowedBookStats() (*domain.BorrowedBookStats, error)
}
