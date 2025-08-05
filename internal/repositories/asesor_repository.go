package repositories

import (
	"lsp-api/internal/models"

	"gorm.io/gorm"
)

type AsesorRepository interface {
	Create(asesor *models.Asesor) error
	Update(asesor *models.Asesor) error
	Delete(id uint) error
	FindByID(id uint) (*models.Asesor, error)
	FindAll() ([]models.Asesor, error)
	FindByNoRegistrasi(noRegistrasi string) (*models.Asesor, error)
}

type asesorRepository struct {
	db *gorm.DB
}

func NewAsesorRepository(db *gorm.DB) AsesorRepository {
	return &asesorRepository{db: db}
}

func (r *asesorRepository) Create(asesor *models.Asesor) error {
	return r.db.Create(asesor).Error
}

func (r *asesorRepository) Update(asesor *models.Asesor) error {
	return r.db.Save(asesor).Error
}

func (r *asesorRepository) Delete(id uint) error {
	return r.db.Delete(&models.Asesor{}, id).Error
}

func (r *asesorRepository) FindByID(id uint) (*models.Asesor, error) {
	var asesor models.Asesor
	err := r.db.Preload("Kompetensi").First(&asesor, id).Error
	if err != nil {
		return nil, err
	}
	return &asesor, nil
}

func (r *asesorRepository) FindAll() ([]models.Asesor, error) {
	var asesors []models.Asesor
	err := r.db.Preload("Kompetensi").Find(&asesors).Error
	if err != nil {
		return nil, err
	}
	return asesors, nil
}

func (r *asesorRepository) FindByNoRegistrasi(noRegistrasi string) (*models.Asesor, error) {
	var asesor models.Asesor
	err := r.db.Preload("Kompetensi").Where("no_registrasi = ?", noRegistrasi).First(&asesor).Error
	if err != nil {
		return nil, err
	}
	return &asesor, nil
}