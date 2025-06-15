package http

import (
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	// setup Sentry
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.SENTRY_DSN,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
		ServerName:       config.APP_ENV + "-lms",
	}); err != nil {
		logrus.Error("Sentry initialization failed:", "error", err)
	}
	router := gin.Default()

	v1 := router.Group("/api/v1/lms")

	// setup Swagger
	docs.SwaggerInfo.Host = config.HOST_PATH
	v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	// Set Middileware
	v1.Use(extractUserInfo())

	user := v1.Group("/user")
	{
		user.POST("", handler.CreateUser)
		user.GET("", handler.GetUser)
		user.PUT("", handler.UpdateUser)
	}

	return &Router{
		router,
	}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
