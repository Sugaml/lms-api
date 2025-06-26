// domain/user.go
package domain

import (
	"errors"
	"time"
)

type User struct {
	BaseModel
	Username  string `gorm:"unique;not null" json:"username"`
	Password  string `gorm:"not null" json:"password"`
	Role      string `gorm:"not null" json:"role"` // student/librarian
	Email     string `gorm:"not null" json:"email"`
	FullName  string `gorm:"column:full_name;not null" json:"full_name"`
	Program   string `json:"program"`
	StudentID string `gorm:"column:student_id" json:"student_id"`
	IsActive  bool   `gorm:"column:is_active;default:false" json:"is_active"`
}

type UserRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	Program   string `json:"program"`
	StudentID string `json:"student_id"`
}

type UserListRequest struct {
	ListRequest
	Username  string `form:"username"`
	Password  string `form:"password"`
	Role      string `form:"role"`
	Email     string `form:"email"`
	FullName  string `form:"full_name"`
	Program   string `form:"program"`
	StudentID string `form:"student_id"`
}

type UserAllUpdateRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	Program   string `json:"program"`
	StudentID string `json:"student_id"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginUserResponse struct {
	AccessToken string `json:"access_token"`
	User        *UserResponse
}

type UserUpdateRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	Program   string `json:"program"`
	StudentID string `json:"student_id"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Program   string    `json:"program"`
	StudentID string    `json:"student_id"`
	IsActive  bool      `json:"is_active"`
}

type StudentResponse struct {
	ID            string `json:"id"`
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	StudentID     string `json:"student_id"`
	Program       string `json:"program"`
	BorrowedCount int    `json:"borrowed_count" default:"10"`
	OverdueCount  int    `json:"overdue_count"  default:"10"`
	Fines         int    `json:"fines"  default:"10"`
	Status        string `json:"status"  default:"clearance"`
	ProfileImage  string `json:"profile_image"`
}

func (u *UserRequest) Validate() error {
	if u.Username == "" {
		return errors.New("username is required")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}
	if u.Role == "" {
		return errors.New("role is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.FullName == "" {
		return errors.New("full name is required")
	}
	if u.Role == "Student" {
		if u.Program == "" {
			return errors.New("program is required")
		}
		if u.StudentID == "" {
			return errors.New("student id is required")
		}
	}
	return nil
}

func (r *UserUpdateRequest) NewUpdate() Map {
	if r.Username != "" {
		return Map{"username": r.Username}
	}
	return nil
}
