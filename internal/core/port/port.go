package port

type Repository interface {
	AuditLogRepository
	UserRepository
	BookRepository
	FineRepository
	BorrowRepository
	NotificationRepository
}
type Service interface {
	AuditLogService
	UserService
	BookService
	FineService
	BorrowService
	NotificationService
}
