package domain

type Notification struct {
	BaseModel
	UserID      int    `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Type        string `gorm:"not null"` // 'due_reminder' | 'fine' | 'request_approved' | 'general'
	IsRead      bool   `gorm:"column:is_read;default:false"`
}

type NotificationRequest struct {
	BaseModel
	UserID      int    `json:"user_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Type        string `json:"type" validate:"required"` // 'due_reminder' | 'fine' | 'request_approved' | 'general'
	IsRead      bool   `json:"is_read"`
}

type UpdateNotificationRequest struct {
	BaseModel
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"` // 'due_reminder' | 'fine' | 'request_approved' | 'general'
	IsRead      bool   `json:"is_read"`
}

type ListNotificationRequest struct {
	ListRequest
	UserID      int    `form:"user_id"`
	Title       string `form:"title"`
	Description string `form:"description"`
	Type        string `form:"type"` // 'due_reminder' | 'fine' | 'request_approved' | 'general'
	IsRead      bool   `form:"is_read"`
}

type NotificationResponse struct {
	BaseModel
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description" `
	Type        string `json:"type"` // 'due_reminder' | 'fine' | 'request_approved' | 'general'
	IsRead      bool   `json:"is_read"`
}

func (u *NotificationRequest) Validate() error {
	return nil
}

func (r *UpdateNotificationRequest) NewUpdate() Map {
	return nil
}
