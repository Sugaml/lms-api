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
