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
		&models.Asesi{},
		&models.Jurusan{},
		&models.Kompetensi{},
		&models.Asesmen{},
		&models.FormAPL01{},
		&models.FormAPL02{},
	)

	if err != nil {
		log.Printf("Error running migrations: %v\n", err)
		return err
	}

	log.Println("Migrations completed successfully")
	return nil
}