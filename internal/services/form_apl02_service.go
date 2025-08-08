package services

import (
	"errors"
	"fmt"
	"time"

	"lsp-api/internal/models"
	"lsp-api/internal/repositories"

	"gorm.io/gorm"
)

type FormAPL02Service interface {
	CreateFormAPL02(idAsesmen uint, kodeUnit, judulUnit, elemen, kuk, status, bukti string) (*models.FormAPL02, error)
	UpdateFormAPL02(id uint, kodeUnit, judulUnit, elemen, kuk, status, bukti string) (*models.FormAPL02, error)
	DeleteFormAPL02(id uint) error
	GetFormAPL02ByID(id uint) (*models.FormAPL02, error)
	GetFormAPL02ByAsesmenID(asesmenID uint) ([]models.FormAPL02, error)
	GetFormAPL02ByAsesiID(asesiID uint) ([]models.FormAPL02, error)
}

type formAPL02Service struct {
	formRepo    repositories.FormAPL02Repository
	asesmenRepo repositories.AsesmenRepository
}

func NewFormAPL02Service(
	formRepo repositories.FormAPL02Repository,
	asesmenRepo repositories.AsesmenRepository,
) FormAPL02Service {
	return &formAPL02Service{
		formRepo:    formRepo,
		asesmenRepo: asesmenRepo,
	}
}

func (s *formAPL02Service) CreateFormAPL02(idAsesmen uint, kodeUnit, judulUnit, elemen, kuk, status, bukti string) (*models.FormAPL02, error) {
	// Verify asesmen exists
	_, err := s.asesmenRepo.FindByID(idAsesmen)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("asesmen not found")
		}
		return nil, err
	}

	// Create new form
	form := &models.FormAPL02{
		IDAsesmen: &idAsesmen,
		KodeUnit:  kodeUnit,
		JudulUnit: judulUnit,
		Elemen:    elemen,
		KUK:       kuk,
		Status:    status,
		Bukti:     bukti,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.formRepo.Create(form)
	if err != nil {
		return nil, fmt.Errorf("failed to create form APL02: %w", err)
	}

	return form, nil
}

func (s *formAPL02Service) UpdateFormAPL02(id uint, kodeUnit, judulUnit, elemen, kuk, status, bukti string) (*models.FormAPL02, error) {
	// Check if form exists
	form, err := s.formRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("form APL02 not found: %w", err)
	}

	// Update form
	if kodeUnit != "" {
		form.KodeUnit = kodeUnit
	}
	if judulUnit != "" {
		form.JudulUnit = judulUnit
	}
	if elemen != "" {
		form.Elemen = elemen
	}
	if kuk != "" {
		form.KUK = kuk
	}
	if status != "" {
		form.Status = status
	}
	if bukti != "" {
		form.Bukti = bukti
	}

	form.UpdatedAt = time.Now()

	err = s.formRepo.Update(form)
	if err != nil {
		return nil, fmt.Errorf("failed to update form APL02: %w", err)
	}

	return form, nil
}

func (s *formAPL02Service) DeleteFormAPL02(id uint) error {
	// Check if form exists
	_, err := s.formRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("form APL02 not found: %w", err)
	}

	return s.formRepo.Delete(id)
}

func (s *formAPL02Service) GetFormAPL02ByID(id uint) (*models.FormAPL02, error) {
	return s.formRepo.FindByID(id)
}

func (s *formAPL02Service) GetFormAPL02ByAsesmenID(asesmenID uint) ([]models.FormAPL02, error) {
	return s.formRepo.FindByAsesmenID(asesmenID)
}

func (s *formAPL02Service) GetFormAPL02ByAsesiID(asesiID uint) ([]models.FormAPL02, error) {
	return s.formRepo.FindByAsesiID(asesiID)
}