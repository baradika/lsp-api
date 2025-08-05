package migrations

import (
	"log"

	"lsp-api/internal/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	log.Println("Running database migrations...")

	// Auto migrate tables
	err := db.AutoMigrate(
		&models.User{},
		&models.Asesor{},
		&models.Kompetensi{},
	)

	if err != nil {
		log.Printf("Error running migrations: %v\n", err)
		return err
	}

	log.Println("Migrations completed successfully")
	return nil
}