package repositories

import (
	"lsp-api/internal/models"

	"gorm.io/gorm"
)

type FormAPL01Repository interface {
	Create(form *models.FormAPL01) error
	Update(form *models.FormAPL01) error
	Delete(id uint) error
	FindByID(id uint) (*models.FormAPL01, error)
	FindByAsesmenID(asesmenID uint) (*models.FormAPL01, error)
	FindByAsesiID(asesiID uint) ([]models.FormAPL01, error)
}

type formAPL01Repository struct {
	db *gorm.DB
}

func NewFormAPL01Repository(db *gorm.DB) FormAPL01Repository {
	return &formAPL01Repository{db: db}
}

func (r *formAPL01Repository) Create(form *models.FormAPL01) error {
	return r.db.Create(form).Error
}

func (r *formAPL01Repository) Update(form *models.FormAPL01) error {
	return r.db.Save(form).Error
}

func (r *formAPL01Repository) Delete(id uint) error {
	return r.db.Delete(&models.FormAPL01{}, id).Error
}

func (r *formAPL01Repository) FindByID(id uint) (*models.FormAPL01, error) {
	var form models.FormAPL01
	err := r.db.Preload("Asesmen").First(&form, id).Error
	if err != nil {
		return nil, err
	}
	return &form, nil
}

func (r *formAPL01Repository) FindByAsesmenID(asesmenID uint) (*models.FormAPL01, error) {
	var form models.FormAPL01
	err := r.db.Where("id_asesmen = ?", asesmenID).First(&form).Error
	if err != nil {
		return nil, err
	}
	return &form, nil
}

func (r *formAPL01Repository) FindByAsesiID(asesiID uint) ([]models.FormAPL01, error) {
	var forms []models.FormAPL01
	err := r.db.Joins("JOIN asesmen ON form_apl01.id_asesmen = asesmen.id_asesmen").Where("asesmen.id_asesi = ?", asesiID).Find(&forms).Error
	if err != nil {
		return nil, err
	}
	return forms, nil
}