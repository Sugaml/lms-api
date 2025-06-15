package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sugaml/lms-api/internal/adaptor/config"
	"github.com/sugaml/lms-api/internal/core/auth"
	"github.com/sugaml/lms-api/internal/core/port"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

type Handler struct {
	svc        port.Service
	config     config.Config
	tokenMaker auth.Maker
}

// NewHandler creates a new Handler instance
func NewHandler(svc port.Service, config config.Config, tokenMaker auth.Maker) *Handler {
	return &Handler{
		svc,
		config,
		tokenMaker,
	}
}

// Router is a wrapper for HTTP router
type Router struct {
	*gin.Engine
}

// NewRouter creates a new HTTP router
func NewRouter(config config.Config, handler Handler) (*Router, error) {
	// Disable debug mode in production
	if config.APP_ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.Use(CORSMiddleware())
	v1 := router.Group("/api/v1/lms")
	// setup Swagger
	docs.SwaggerInfo.Host = config.HOST_PATH
	v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	user := v1.Group("/users")
	{
		user.POST("", handler.CreateUser)
		user.POST("/login", handler.LoginUser)
	}

	v1.Use(authMiddleware(handler.tokenMaker))

	userAuth := v1.Group("/users")
	{
		userAuth.GET("", handler.ListUser)
		userAuth.GET("/:id", handler.GetUser)
		userAuth.PUT("", handler.UpdateUser)
		userAuth.DELETE("/:id", handler.DeleteUser)
	}

	category := v1.Group("/categories")
	{
		category.POST("", handler.CreateCategory)
		category.GET("", handler.ListCategory)
		category.GET("/:id", handler.GetCategory)
		category.PUT("/:id", handler.UpdateCategory)
		category.DELETE("/:id", handler.DeleteCategory)
	}

	program := v1.Group("/programs")
	{
		program.POST("", handler.CreateProgram)
		program.GET("", handler.ListProgram)
		program.GET("/:id", handler.GetProgram)
		program.PUT("/:id", handler.UpdateProgram)
		program.DELETE("/:id", handler.DeleteProgram)
	}

	auditlog := v1.Group("/auditlog")
	{
		auditlog.POST("", handler.CreateAuditLog)
		auditlog.GET("", handler.ListAuditLog)
		auditlog.GET("/:id", handler.GetAuditLog)
		auditlog.PUT("/:id", handler.UpdateAuditLog)
		auditlog.DELETE("/:id", handler.DeleteAuditLog)
	}

	book := v1.Group("/books")
	{
		book.POST("", handler.CreateBook)
		book.GET("", handler.ListBook)
		book.GET("/:id", handler.GetBook)
		book.PUT("", handler.UpdateBook)
		book.DELETE("/:id", handler.DeleteBook)
	}

	student := v1.Group("/students")
	{
		student.POST("", handler.CreateStudent)
		student.GET("", handler.ListStudent)
	}

	report := v1.Group("/reports")
	{
		report.GET("dashboard-stats", handler.GetLibraryDashboardStats)
		report.GET("borrowedbookstats", handler.GetBorrowedBookStats)
	}
	borrow := v1.Group("/borrows")
	{
		borrow.POST("", handler.CreateBorrow)
		borrow.GET("", handler.ListBorrow)
		borrow.GET("/:id", handler.GetBorrow)
		borrow.PUT("/:id", handler.UpdateBorrow)
		borrow.DELETE("/:id", handler.DeleteBorrow)
	}

	fine := v1.Group("/fines")
	{
		fine.POST("", handler.CreateFine)
		fine.GET("", handler.ListFine)
		fine.GET("/:id", handler.GetFine)
		fine.PUT("", handler.UpdateFine)
		fine.DELETE("/:id", handler.DeleteFine)
	}

	notification := v1.Group("/notifications")
	{
		notification.POST("", handler.CreateNotification)
		notification.GET("", handler.ListNotification)
		notification.GET("/:id", handler.GetNotification)
		notification.PUT("", handler.UpdateNotification)
		notification.DELETE("/:id", handler.DeleteNotification)
	}

	return &Router{
		router,
	}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
