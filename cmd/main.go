package main

import (
	"fmt"
	"log"

	"lsp-api/internal/config"
	"lsp-api/internal/controllers"
	"lsp-api/internal/middleware"
	"lsp-api/internal/repositories"
	"lsp-api/internal/services"
	"lsp-api/migrations"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Run migrations
	err = migrations.RunMigrations(db)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	asesorRepo := repositories.NewAsesorRepository(db)
	kompetensiRepo := repositories.NewKompetensiRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo, cfg)
	asesorService := services.NewAsesorService(asesorRepo, kompetensiRepo)

	// Initialize controllers
	authController := controllers.NewAuthController(authService)
	asesorController := controllers.NewAsesorController(asesorService)

	// Initialize middleware
	authMiddleware := middleware.AuthMiddleware(authService)

	// Initialize router
	router := gin.Default()

	// API routes
	apiV1 := router.Group("/api/v1")
	{
		// Register auth routes
		authController.RegisterRoutes(apiV1, authMiddleware)

		// Register asesor routes
		asesorController.RegisterRoutes(apiV1, authMiddleware)
	}

	// Start server
	serverAddr := fmt.Sprintf(":%s", cfg.AppPort)
	log.Printf("Server starting on %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}