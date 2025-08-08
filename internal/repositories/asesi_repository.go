package repositories

import (
    "lsp-api/internal/models"
    "gorm.io/gorm"
)

type AsesiRepository interface {
    Create(asesi *models.Asesi) error
    Update(asesi *models.Asesi) error
    Delete(id uint) error
    FindByID(id uint) (*models.Asesi, error)
    FindByUserID(userID uint) (*models.Asesi, error)
    FindAll() ([]*models.Asesi, error)
}

type asesiRepository struct {
    db *gorm.DB
}

func NewAsesiRepository(db *gorm.DB) AsesiRepository {
    return &asesiRepository{db: db}
}

func (r *asesiRepository) Create(asesi *models.Asesi) error {
    return r.db.Create(asesi).Error
}

func (r *asesiRepository) Update(asesi *models.Asesi) error {
    return r.db.Save(asesi).Error
}

func (r *asesiRepository) Delete(id uint) error {
    return r.db.Delete(&models.Asesi{}, id).Error
}

func (r *asesiRepository) FindByID(id uint) (*models.Asesi, error) {
    var asesi models.Asesi
    err := r.db.Preload("User").Preload("Jurusan").First(&asesi, id).Error
    if err != nil {
        return nil, err
    }
    return &asesi, nil
}

func (r *asesiRepository) FindByUserID(userID uint) (*models.Asesi, error) {
    var asesi models.Asesi
    err := r.db.Preload("User").Preload("Jurusan").Where("id_user = ?", userID).First(&asesi).Error
    if err != nil {
        return nil, err
    }
    return &asesi, nil
}

func (r *asesiRepository) FindAll() ([]*models.Asesi, error) {
    var asesis []*models.Asesi
    err := r.db.Preload("User").Preload("Jurusan").Find(&asesis).Error
    return asesis, err
}