package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sugaml/lms-api/internal/adaptor/config"
	"github.com/sugaml/lms-api/internal/core/port"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

type Handler struct {
	svc    port.Service
	config config.Config
}

// NewHandler creates a new Handler instance
func NewHandler(svc port.Service, config config.Config) *Handler {
	return &Handler{
		svc,
		config,
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

	v1 := router.Group("/api/v1/lms")

	// setup Swagger
	docs.SwaggerInfo.Host = config.HOST_PATH
	v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	// Set Middileware
	v1.Use(extractUserInfo())

	user := v1.Group("/users")
	{
		user.POST("", handler.CreateUser)
		user.GET("", handler.ListUser)
		user.GET("/:id", handler.GetUser)
		user.PUT("", handler.UpdateUser)
		user.DELETE("/:id", handler.DeleteUser)
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

	return &Router{
		router,
	}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
