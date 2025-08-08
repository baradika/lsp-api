package repositories

import (
	"lsp-api/internal/models"

	"gorm.io/gorm"
)

type FormAPL02Repository interface {
	Create(form *models.FormAPL02) error
	Update(form *models.FormAPL02) error
	Delete(id uint) error
	FindByID(id uint) (*models.FormAPL02, error)
	FindByAsesmenID(asesmenID uint) ([]models.FormAPL02, error)
	FindByAsesiID(asesiID uint) ([]models.FormAPL02, error)
}

type formAPL02Repository struct {
	db *gorm.DB
}

func NewFormAPL02Repository(db *gorm.DB) FormAPL02Repository {
	return &formAPL02Repository{db: db}
}

func (r *formAPL02Repository) Create(form *models.FormAPL02) error {
	return r.db.Create(form).Error
}

func (r *formAPL02Repository) Update(form *models.FormAPL02) error {
	return r.db.Save(form).Error
}

func (r *formAPL02Repository) Delete(id uint) error {
	return r.db.Delete(&models.FormAPL02{}, id).Error
}

func (r *formAPL02Repository) FindByID(id uint) (*models.FormAPL02, error) {
	var form models.FormAPL02
	err := r.db.Preload("Asesmen").First(&form, id).Error
	if err != nil {
		return nil, err
	}
	return &form, nil
}

func (r *formAPL02Repository) FindByAsesmenID(asesmenID uint) ([]models.FormAPL02, error) {
	var forms []models.FormAPL02
	err := r.db.Where("id_asesmen = ?", asesmenID).Find(&forms).Error
	if err != nil {
		return nil, err
	}
	return forms, nil
}

func (r *formAPL02Repository) FindByAsesiID(asesiID uint) ([]models.FormAPL02, error) {
	var forms []models.FormAPL02
	err := r.db.Joins("JOIN asesmen ON form_apl02.id_asesmen = asesmen.id_asesmen").Where("asesmen.id_asesi = ?", asesiID).Find(&forms).Error
	if err != nil {
		return nil, err
	}
	return forms, nil
}