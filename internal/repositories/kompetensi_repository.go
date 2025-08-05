package repositories

import (
	"lsp-api/internal/models"

	"gorm.io/gorm"
)

type KompetensiRepository interface {
	Create(kompetensi *models.Kompetensi) error
	FindByID(id uint) (*models.Kompetensi, error)
	FindAll() ([]models.Kompetensi, error)
	FindByIDs(ids []uint) ([]models.Kompetensi, error)
}

type kompetensiRepository struct {
	db *gorm.DB
}

func NewKompetensiRepository(db *gorm.DB) KompetensiRepository {
	return &kompetensiRepository{db: db}
}

func (r *kompetensiRepository) Create(kompetensi *models.Kompetensi) error {
	return r.db.Create(kompetensi).Error
}

func (r *kompetensiRepository) FindByID(id uint) (*models.Kompetensi, error) {
	var kompetensi models.Kompetensi
	err := r.db.First(&kompetensi, id).Error
	if err != nil {
		return nil, err
	}
	return &kompetensi, nil
}

func (r *kompetensiRepository) FindAll() ([]models.Kompetensi, error) {
	var kompetensi []models.Kompetensi
	err := r.db.Find(&kompetensi).Error
	if err != nil {
		return nil, err
	}
	return kompetensi, nil
}

func (r *kompetensiRepository) FindByIDs(ids []uint) ([]models.Kompetensi, error) {
	var kompetensi []models.Kompetensi
	err := r.db.Where("id IN ?", ids).Find(&kompetensi).Error
	if err != nil {
		return nil, err
	}
	return kompetensi, nil
}