package port

type Repository interface {
	AuditLogRepository
	UserRepository
	BookRepository
}
type Service interface {
	AuditLogService
	UserService
	BookService
}
