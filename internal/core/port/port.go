package port

type Repository interface {
	AuditLogRepository
	UserRepository
}
type Service interface {
	AuditLogService
	UserService
}
