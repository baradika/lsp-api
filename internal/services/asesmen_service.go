package services

import (
	"errors"
	"fmt"
	"time"

	"lsp-api/internal/models"
	"lsp-api/internal/repositories"

	"gorm.io/gorm"
)

type AsesmenService interface {
	CreateAsesmen(idAsesi, idAsesor, idSkema uint, tuk string) (*models.Asesmen, error)
	UpdateAsesmen(id uint, status, hasil, catatan string) (*models.Asesmen, error)
	DeleteAsesmen(id uint) error
	GetAsesmenByID(id uint) (*models.Asesmen, error)
	GetAsesmenByAsesiID(asesiID uint) ([]models.Asesmen, error)
	GetAsesmenByAsesorID(asesorID uint) ([]models.Asesmen, error)
}

type asesmenService struct {
	asesmenRepo repositories.AsesmenRepository
	asesiRepo   repositories.AsesiRepository
	asesorRepo  repositories.AsesorRepository
}

func NewAsesmenService(
	asesmenRepo repositories.AsesmenRepository,
	asesiRepo repositories.AsesiRepository,
	asesorRepo repositories.AsesorRepository,
) AsesmenService {
	return &asesmenService{
		asesmenRepo: asesmenRepo,
		asesiRepo:   asesiRepo,
		asesorRepo:  asesorRepo,
	}
}

func (s *asesmenService) CreateAsesmen(idAsesi, idAsesor, idSkema uint, tuk string) (*models.Asesmen, error) {
	// Verify asesi exists
	_, err := s.asesiRepo.FindByID(idAsesi)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("asesi not found")
		}
		return nil, err
	}

	// Verify asesor exists
	_, err = s.asesorRepo.FindByID(idAsesor)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("asesor not found")
		}
		return nil, err
	}

	// Create new asesmen
	now := time.Now()
	asesmen := &models.Asesmen{
		IDAsesi:      &idAsesi,
		IDAsesor:     &idAsesor,
		IDSkema:      &idSkema,
		TUK:          tuk,
		Status:       "Draft",
		Hasil:        "Belum Selesai",
		TanggalMulai: &now,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	err = s.asesmenRepo.Create(asesmen)
	if err != nil {
		return nil, fmt.Errorf("failed to create asesmen: %w", err)
	}

	return asesmen, nil
}

func (s *asesmenService) UpdateAsesmen(id uint, status, hasil, catatan string) (*models.Asesmen, error) {
	// Check if asesmen exists
	asesmen, err := s.asesmenRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("asesmen not found: %w", err)
	}

	// Update asesmen
	if status != "" {
		asesmen.Status = status
	}
	if hasil != "" {
		asesmen.Hasil = hasil
	}
	if catatan != "" {
		asesmen.Catatan = catatan
	}

	if status == "Selesai" {
		now := time.Now()
		asesmen.TanggalSelesai = &now
	}

	asesmen.UpdatedAt = time.Now()

	err = s.asesmenRepo.Update(asesmen)
	if err != nil {
		return nil, fmt.Errorf("failed to update asesmen: %w", err)
	}

	return asesmen, nil
}

func (s *asesmenService) DeleteAsesmen(id uint) error {
	// Check if asesmen exists
	_, err := s.asesmenRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("asesmen not found: %w", err)
	}

	return s.asesmenRepo.Delete(id)
}

func (s *asesmenService) GetAsesmenByID(id uint) (*models.Asesmen, error) {
	return s.asesmenRepo.FindByID(id)
}

func (s *asesmenService) GetAsesmenByAsesiID(asesiID uint) ([]models.Asesmen, error) {
	return s.asesmenRepo.FindByAsesiID(asesiID)
}

func (s *asesmenService) GetAsesmenByAsesorID(asesorID uint) ([]models.Asesmen, error) {
	return s.asesmenRepo.FindByAsesorID(asesorID)
}