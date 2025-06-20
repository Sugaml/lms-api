package repository

import (
	"time"

	"github.com/sugaml/lms-api/internal/core/domain"
)

func (r *Repository) GetLibraryDashboardStats() (*domain.LibraryDashboardStats, error) {
	var stats domain.LibraryDashboardStats
	now := time.Now()

	//Count total students
	if err := r.db.Model(&domain.User{}).
		Where("role = ?", "student").
		Count(&stats.TotalStudents).Error; err != nil {
		return nil, err
	}

	//Count active students
	if err := r.db.Model(&domain.User{}).
		Where("role = ? AND is_active = ?", "student", true).
		Count(&stats.ActiveStudents).Error; err != nil {
		return nil, err
	}

	//Count total active books
	if err := r.db.Model(&domain.Book{}).
		Count(&stats.TotalBooks).Error; err != nil {
		return nil, err
	}

	// Count total pending books
	if err := r.db.Model(&domain.BorrowedBook{}).
		Where("status = ? AND returned_date IS NULL AND is_active = ?", "borrowed", true).
		Count(&stats.PendingRequests).Error; err != nil {
		return nil, err
	}

	// Count total borrowed books
	if err := r.db.Model(&domain.BorrowedBook{}).
		Where("status = ? AND due_date < ? AND returned_date IS NULL AND is_active = ?", "borrowed", now, true).
		Count(&stats.BorrowedBooks).Error; err != nil {
		return nil, err
	}

	// Count overdue books
	if err := r.db.Model(&domain.BorrowedBook{}).
		Where("status = ? AND due_date < ? AND returned_date IS NULL AND is_active = ?", "borrowed", now, true).
		Count(&stats.OverdueBooks).Error; err != nil {
		return nil, err
	}

	return &stats, nil
}

func (r *Repository) GetBorrowedBookStats() (*domain.BorrowedBookStats, error) {
	var stats domain.BorrowedBookStats
	now := time.Now()

	// Count total borrowed books
	if err := r.db.Model(&domain.BorrowedBook{}).
		Count(&stats.TotalBorrowedBooks).Error; err != nil {
		return nil, err
	}

	// Count overdue books
	if err := r.db.Model(&domain.BorrowedBook{}).
		Where("status = ? AND due_date < ? AND returned_date IS NULL AND is_active = ?", "borrowed", now, true).
		Count(&stats.TotalOverdueBooks).Error; err != nil {
		return nil, err
	}

	// Count pending requests
	if err := r.db.Model(&domain.BorrowedBook{}).
		Where("status = ? AND is_active = ?", "pending", true).
		Count(&stats.PendingRequests).Error; err != nil {
		return nil, err
	}

	// Count due soon (within 3 days)
	threeDaysLater := now.Add(72 * time.Hour)
	if err := r.db.Model(&domain.BorrowedBook{}).
		Where("status = ? AND due_date BETWEEN ? AND ? AND is_active = ?", "borrowed", now, threeDaysLater, true).
		Count(&stats.DueSoon).Error; err != nil {
		return nil, err
	}

	return &stats, nil
}

func (r *Repository) GetBookProgramstats() (*[]domain.BookProgramstats, error) {
	var stats []domain.BookProgramstats
	if err := r.db.Model(&domain.Book{}).
		Select("program, count(*) as count").
		Group("program").
		Find(&stats).Error; err != nil {
		return nil, err
	}
	return &stats, nil
}
