package port

type Repository interface {
	AuditLogRepository
	UserRepository
	CategoryRepository
	ProgramRepository
	BookRepository
	FineRepository
	BorrowRepository
	NotificationRepository
}
type Service interface {
	AuditLogService
	UserService
	CategoryService
	ProgramService
	BookService
	FineService
	BorrowService
	NotificationService
}
