package port

type Repository interface {
	AuditLogRepository
	UserRepository
	CategoryRepository
	BookRepository
	FineRepository
	BorrowRepository
	NotificationRepository
}
type Service interface {
	AuditLogService
	UserService
	CategoryService
	BookService
	FineService
	BorrowService
	NotificationService
}
