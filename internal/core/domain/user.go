// domain/user.go
package domain

import "time"

type User struct {
	BaseModel
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"not null"`
	Email     string `gorm:"not null"`
	FullName  string `gorm:"column:full_name;not null"`
	Program   string
	StudentID string `gorm:"column:student_id"`
}

type UserRequest struct {
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"not null"`
	Email     string `gorm:"not null"`
	FullName  string `gorm:"column:full_name;not null"`
	Program   string
	StudentID string `gorm:"column:student_id"`
}

type UserListRequest struct {
	ListRequest
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"not null"`
	Email     string `gorm:"not null"`
	FullName  string `gorm:"column:full_name;not null"`
	Program   string
	StudentID string `gorm:"column:student_id"`
}

type UserAllUpdateRequest struct {
	Username  string `json:"username"`
	Password  string
	Role      string
	Email     string
	FullName  string
	Program   string
	StudentID string
}

type UserUpdateRequest struct {
	Username  string `json:"username"`
	Password  string
	Role      string
	Email     string
	FullName  string
	Program   string
	StudentID string
}

func (r *UserUpdateRequest) NewUpdate() Map {
	return nil
}

type UserResponse struct {
	Username  string
	Password  string
	Role      string
	Email     string
	FullName  string
	Program   string
	StudentID string
}

func (u *UserRequest) Validate() error {
	return nil
}

type BorrowedBook struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`
	UserID       int       `gorm:"not null"`
	BookID       int       `gorm:"not null"`
	BorrowedDate time.Time `gorm:"column:borrowed_date;autoCreateTime"`
	DueDate      time.Time `gorm:"not null"`
	ReturnedDate *time.Time
	RenewalCount int    `gorm:"default:0"`
	Status       string `gorm:"not null"` // 'borrowed' | 'returned' | 'overdue'
}

type Fine struct {
	ID             int        `gorm:"primaryKey;autoIncrement"`
	UserID         int        `gorm:"not null"`
	BorrowedBookID int        `gorm:"column:borrowed_book_id;not null"`
	Amount         int        `gorm:"not null"` // in paisa
	Reason         string     `gorm:"not null"`
	Status         string     `gorm:"not null"` // 'pending' | 'paid'
	CreatedAt      time.Time  `gorm:"column:created_at;autoCreateTime"`
	PaidAt         *time.Time `gorm:"column:paid_at"`
}

type BookRequest struct {
	ID           int        `gorm:"primaryKey;autoIncrement"`
	UserID       int        `gorm:"not null"`
	BookID       int        `gorm:"not null"`
	Status       string     `gorm:"not null"` // 'pending' | 'approved' | 'rejected'
	RequestDate  time.Time  `gorm:"column:request_date;autoCreateTime"`
	ApprovedDate *time.Time `gorm:"column:approved_date"`
	ApprovedBy   *int       `gorm:"column:approved_by"`
}

type Notification struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	UserID      int       `gorm:"not null"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Type        string    `gorm:"not null"` // 'due_reminder' | 'fine' | 'request_approved' | 'general'
	IsRead      bool      `gorm:"column:is_read;default:false"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
}
