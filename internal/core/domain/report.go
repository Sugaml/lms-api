package domain

type LibraryDashboardStats struct {
	ActiveStudents  int64 `json:"activeStudents"`
	AvailableBooks  int64 `json:"availableBooks"`
	BorrowedBooks   int64 `json:"borrowedBooks"`
	OverdueBooks    int64 `json:"overdueBooks"`
	PendingRequests int64 `json:"pendingRequests"`
	TotalBooks      int64 `json:"totalBooks"`
	TotalFines      int64 `json:"totalFines"`
	TotalStudents   int64 `json:"totalStudents"`
}

type BorrowedBookStats struct {
	TotalBorrowedBooks int64 `json:"totalBorrowedBooks"`
	TotalOverdueBooks  int64 `json:"totalOverdueBooks"`
	PendingRequests    int64 `json:"pendingRequests"`
	DueSoon            int64 `json:"dueSoon"`
}

type BookProgramstats struct {
	Program string `json:"program"`
	Count   int64  `json:"count"`
}

type InventoryStats struct {
	TotalBooks      int64 `json:"totalBooks"`
	AvailableBooks  int64 `json:"availableBooks"`
	BorrowedBooks   int64 `json:"borrowedBooks"`
	OverdueBooks    int64 `json:"overdueBooks"`
	TotalStudents   int64 `json:"totalStudents"`
	ActiveStudents  int64 `json:"activeStudents"`
	PendingRequests int64 `json:"pendingRequests"`
	TotalFines      int64 `json:"totalFines"`
}
