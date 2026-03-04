package router

import (
	"go-crud/internal/domain/user"
	"go-crud/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupUserRoutes(db *sqlx.DB) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	api := r.Group("/api/v1")

	// User domain
	userRepo := user.NewRepository(db)
	userSvc := user.NewService(userRepo)
	userHandler := user.NewHandler(userSvc)
	userHandler.RegisterRoutes(api)

	return r
}
