package port

import "github.com/sugaml/lms-api/internal/core/domain"

type ReportRepository interface {
	GetLibraryDashboardStats() (*domain.LibraryDashboardStats, error)
	GetBorrowedBookStats() (*domain.BorrowedBookStats, error)
	GetBookProgramstats() (*[]domain.BookProgramstats, error)
	GetInventorystats() (*domain.InventoryStats, error)
}

type ReportService interface {
	GetLibraryDashboardStats() (*domain.LibraryDashboardStats, error)
	GetBorrowedBookStats() (*domain.BorrowedBookStats, error)
	GetBookProgramstats() (*[]domain.BookProgramstats, error)
	GetInventorystats() (*domain.InventoryStats, error)
}
