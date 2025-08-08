package repositories

import (
	"lsp-api/internal/models"

	"gorm.io/gorm"
)

type AsesmenRepository interface {
	Create(asesmen *models.Asesmen) error
	Update(asesmen *models.Asesmen) error
	Delete(id uint) error
	FindByID(id uint) (*models.Asesmen, error)
	FindByAsesiID(asesiID uint) ([]models.Asesmen, error)
	FindByAsesorID(asesorID uint) ([]models.Asesmen, error)
}

type asesmenRepository struct {
	db *gorm.DB
}

func NewAsesmenRepository(db *gorm.DB) AsesmenRepository {
	return &asesmenRepository{db: db}
}

func (r *asesmenRepository) Create(asesmen *models.Asesmen) error {
	return r.db.Create(asesmen).Error
}

func (r *asesmenRepository) Update(asesmen *models.Asesmen) error {
	return r.db.Save(asesmen).Error
}

func (r *asesmenRepository) Delete(id uint) error {
	return r.db.Delete(&models.Asesmen{}, id).Error
}

func (r *asesmenRepository) FindByID(id uint) (*models.Asesmen, error) {
	var asesmen models.Asesmen
	err := r.db.Preload("Asesi").Preload("Asesor").First(&asesmen, id).Error
	if err != nil {
		return nil, err
	}
	return &asesmen, nil
}

func (r *asesmenRepository) FindByAsesiID(asesiID uint) ([]models.Asesmen, error) {
	var asesmens []models.Asesmen
	err := r.db.Where("id_asesi = ?", asesiID).Preload("Asesor").Find(&asesmens).Error
	if err != nil {
		return nil, err
	}
	return asesmens, nil
}

func (r *asesmenRepository) FindByAsesorID(asesorID uint) ([]models.Asesmen, error) {
	var asesmens []models.Asesmen
	err := r.db.Where("id_asesor = ?", asesorID).Preload("Asesi").Find(&asesmens).Error
	if err != nil {
		return nil, err
	}
	return asesmens, nil
}