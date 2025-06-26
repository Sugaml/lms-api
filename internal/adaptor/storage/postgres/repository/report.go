package repository

import (
	"fmt"
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

func (r *Repository) GetMonthlyChartData() ([]domain.ChartData, error) {
	var results []domain.ChartData

	type TempData struct {
		Month         string
		YearMonth     string
		Borrowed      int
		Returned      int
		Due           int
		Requests      int
		TotalStudents int
		BooksAdded    int
	}

	var data []TempData

	err := r.db.Raw(`
			WITH borrow_summary AS (
				SELECT 
					TO_CHAR(borrowed_date, 'Mon') AS month,
					TO_CHAR(borrowed_date, 'YYYY-MM') AS year_month,
					COUNT(*) FILTER (WHERE status = 'borrowed') AS borrowed,
					COUNT(*) FILTER (WHERE status = 'returned') AS returned,
					COUNT(*) FILTER (WHERE status = 'overdue') AS due,
					COUNT(*) FILTER (WHERE status = 'pending') AS requests
				FROM borrowed_books
				GROUP BY year_month, month
			),
			student_summary AS (
				SELECT 
					TO_CHAR(created_at, 'YYYY-MM') AS year_month,
					COUNT(*) AS total_students
				FROM users
				WHERE role = 'student'
				GROUP BY year_month
			),
			book_summary AS (
				SELECT 
					TO_CHAR(created_at, 'YYYY-MM') AS year_month,
					COUNT(*) AS books_added
				FROM books
				GROUP BY year_month
			)
			SELECT 
				bs.month,
				bs.year_month,
				bs.borrowed,
				bs.returned,
				bs.due,
				bs.requests,
				COALESCE(ss.total_students, 0) AS total_students,
				COALESCE(bo.books_added, 0) AS books_added
			FROM borrow_summary bs
			LEFT JOIN student_summary ss ON ss.year_month = bs.year_month
			LEFT JOIN book_summary bo ON bo.year_month = bs.year_month
			ORDER BY bs.year_month ASC
		`).Scan(&data).Error

	if err != nil {
		return nil, err
	}

	for _, d := range data {
		results = append(results, domain.ChartData{
			Month:         d.Month,
			Date:          d.YearMonth,
			Borrowed:      d.Borrowed,
			Returned:      d.Returned,
			Due:           d.Due,
			Requests:      d.Requests,
			TotalStudents: d.TotalStudents,
			BooksAdded:    d.BooksAdded,
		})
	}

	return results, nil
}

func (r *Repository) GetDailyChartData(startDateStr, endDateStr, rangeType string) ([]domain.ChartData, error) {
	var results []domain.ChartData

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return nil, err
	}
	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return nil, err
	}

	type TempData struct {
		Label         string
		GroupKey      string
		Borrowed      int
		Returned      int
		Due           int
		Requests      int
		TotalStudents int
		BooksAdded    int
	}

	// Properly quoted format strings for TO_CHAR()
	var labelFmt, groupFmt string

	switch rangeType {
	case "daily":
		labelFmt = "'Dy DD'"
		groupFmt = "YYYY-MM-DD"
	case "monthly":
		labelFmt = "'Mon'"
		groupFmt = "YYYY-MM"
	case "yearly":
		labelFmt = "'YYYY'"
		groupFmt = "YYYY"
	case "weekly", "":
		fallthrough
	default:
		labelFmt = "'Wk ''IW'" // escaped single quotes for 'Wk 'IW'
		groupFmt = "IYYY-IW"
	}

	query := fmt.Sprintf(`
		WITH borrow_summary AS (
			SELECT 
				TO_CHAR(borrowed_date, %s) AS label,
				TO_CHAR(borrowed_date, '%s') AS group_key,
				COUNT(*) FILTER (WHERE status = 'borrowed') AS borrowed,
				COUNT(*) FILTER (WHERE status = 'returned') AS returned,
				COUNT(*) FILTER (WHERE status = 'overdue') AS due,
				COUNT(*) FILTER (WHERE status = 'pending') AS requests
			FROM borrowed_books
			WHERE borrowed_date BETWEEN ? AND ?
			GROUP BY group_key, label
		),
		student_summary AS (
			SELECT 
				TO_CHAR(created_at, '%s') AS group_key,
				COUNT(*) AS total_students
			FROM users
			WHERE role = 'student' AND created_at BETWEEN ? AND ?
			GROUP BY group_key
		),
		book_summary AS (
			SELECT 
				TO_CHAR(created_at, '%s') AS group_key,
				COUNT(*) AS books_added
			FROM books
			WHERE created_at BETWEEN ? AND ?
			GROUP BY group_key
		)
		SELECT 
			bs.label,
			bs.group_key,
			bs.borrowed,
			bs.returned,
			bs.due,
			bs.requests,
			COALESCE(ss.total_students, 0) AS total_students,
			COALESCE(bo.books_added, 0) AS books_added
		FROM borrow_summary bs
		LEFT JOIN student_summary ss ON ss.group_key = bs.group_key
		LEFT JOIN book_summary bo ON bo.group_key = bs.group_key
		ORDER BY bs.group_key ASC
	`, labelFmt, groupFmt, groupFmt, groupFmt)

	var rawData []TempData

	err = r.db.Raw(query, startDateStr, endDateStr, startDateStr, endDateStr, startDateStr, endDateStr).Scan(&rawData).Error
	if err != nil {
		return nil, err
	}

	// Map group_key => TempData for easy lookup
	dataMap := make(map[string]TempData)
	for _, row := range rawData {
		dataMap[row.GroupKey] = row
	}

	// If daily range, fill missing days with zeroes
	if rangeType == "daily" {
		for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
			key := d.Format("2006-01-02")
			label := d.Format("Mon 02")
			if row, found := dataMap[key]; found {
				results = append(results, domain.ChartData{
					Month:         row.Label,
					Date:          row.GroupKey,
					Borrowed:      row.Borrowed,
					Returned:      row.Returned,
					Due:           row.Due,
					Requests:      row.Requests,
					TotalStudents: row.TotalStudents,
					BooksAdded:    row.BooksAdded,
				})
			} else {
				results = append(results, domain.ChartData{
					Month:         label,
					Date:          key,
					Borrowed:      0,
					Returned:      0,
					Due:           0,
					Requests:      0,
					TotalStudents: 0,
					BooksAdded:    0,
				})
			}
		}
	} else {
		// For non-daily, just return what DB gave
		for _, row := range rawData {
			results = append(results, domain.ChartData{
				Month:         row.Label,
				Date:          row.GroupKey,
				Borrowed:      row.Borrowed,
				Returned:      row.Returned,
				Due:           row.Due,
				Requests:      row.Requests,
				TotalStudents: row.TotalStudents,
				BooksAdded:    row.BooksAdded,
			})
		}
	}

	return results, nil
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

func (r *Repository) GetInventorystats() (*domain.InventoryStats, error) {
	var stats domain.InventoryStats
	var totalBooks int64
	var borrowedBooks int64
	var overdueBooks int64
	var totalStudents int64
	var activeStudents int64
	var pendingRequests int64
	var totalFines int64

	// Queries
	r.db.Model(&domain.Book{}).Count(&totalBooks)
	r.db.Model(&domain.BorrowedBook{}).Where("status = ?", "borrowed").Count(&borrowedBooks)
	r.db.Model(&domain.BorrowedBook{}).Where("status = ?", "overdue").Count(&overdueBooks)
	r.db.Model(&domain.User{}).Where("role = ?", "student").Count(&totalStudents)
	r.db.Model(&domain.User{}).Where("role = ? AND is_active = ?", "student", true).Count(&activeStudents)
	r.db.Model(&domain.BorrowedBook{}).Where("status = ?", "pending").Count(&pendingRequests)
	r.db.Model(&domain.Fine{}).Where("status = ?", "pending").Select("SUM(amount)").Scan(&totalFines)

	// Available books = sum of all book copies - borrowed books
	var availableBooks int64
	r.db.Model(&domain.Book{}).Select("SUM(total_copies)").Scan(&availableBooks)
	availableBooks = availableBooks - borrowedBooks

	stats = domain.InventoryStats{
		TotalBooks:      totalBooks,
		AvailableBooks:  availableBooks,
		BorrowedBooks:   borrowedBooks,
		OverdueBooks:    overdueBooks,
		TotalStudents:   totalStudents,
		ActiveStudents:  activeStudents,
		PendingRequests: pendingRequests,
		TotalFines:      totalFines,
	}

	return &stats, nil
}
