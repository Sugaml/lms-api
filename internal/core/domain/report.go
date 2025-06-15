package domain

type LibraryDashboard struct {
	ActiveStudents  int `json:"activeStudents"`
	AvailableBooks  int `json:"availableBooks"`
	BorrowedBooks   int `json:"borrowedBooks"`
	OverdueBooks    int `json:"overdueBooks"`
	PendingRequests int `json:"pendingRequests"`
	TotalBooks      int `json:"totalBooks"`
	TotalFines      int `json:"totalFines"`
	TotalStudents   int `json:"totalStudents"`
}

type BorrowedBookStats struct {
	TotalBorrowedBooks int64 `json:"totalBorrowedBooks"`
	TotalOverdueBooks  int64 `json:"totalOverdueBooks"`
	PendingRequests    int64 `json:"pendingRequests"`
	DueSoon            int64 `json:"dueSoon"`
}
