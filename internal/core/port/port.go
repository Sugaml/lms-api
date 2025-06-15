package port

type Repository interface {
	AuditLogRepository
	UserRepository
	CategoryRepository
	ProgramRepository
	BookRepository
	FineRepository
	BorrowRepository
	ReportRepository
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
	ReportService
	NotificationService
}
