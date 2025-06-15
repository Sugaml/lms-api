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

type UserUpdateRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	Program   string `json:"program"`
	StudentID string `json:"student_id"`
}

func (r *UserUpdateRequest) NewUpdate() Map {
	return nil
}

type UserResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Program   string    `json:"program"`
	StudentID string    `json:"student_id"`
}

func (u *UserRequest) Validate() error {
	return nil
}

func (r *BookRequest) Validate() error {
	return nil
}
