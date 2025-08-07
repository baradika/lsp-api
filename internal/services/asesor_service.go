package services

import (
	"errors"
	"fmt"

	"lsp-api/internal/models"
	"lsp-api/internal/repositories"

	"gorm.io/gorm"
)

type AsesorService interface {
	CreateAsesor(namaLengkap, noRegistrasi, email, noTelepon string, kompetensiIDs []uint) (*models.Asesor, error)
	UpdateAsesor(id uint, namaLengkap, noRegistrasi, email, noTelepon string, kompetensiIDs []uint) (*models.Asesor, error)
	DeleteAsesor(id uint) error
	GetAsesorByID(id uint) (*models.Asesor, error)
	GetAllAsesors() ([]models.Asesor, error)
	GetAsesorByNoRegistrasi(noRegistrasi string) (*models.Asesor, error)
}

type asesorService struct {
	asesorRepo     repositories.AsesorRepository
	kompetensiRepo repositories.KompetensiRepository
}

func NewAsesorService(asesorRepo repositories.AsesorRepository, kompetensiRepo repositories.KompetensiRepository) AsesorService {
	return &asesorService{
		asesorRepo:     asesorRepo,
		kompetensiRepo: kompetensiRepo,
	}
}

func (s *asesorService) CreateAsesor(namaLengkap, noRegistrasi, email, noTelepon string, kompetensiIDs []uint) (*models.Asesor, error) {
	// Check if asesor with the same registration number already exists
	_, err := s.asesorRepo.FindByNoRegistrasi(noRegistrasi)
	if err == nil {
		return nil, errors.New("asesor with this registration number already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Get kompetensi by IDs
	kompetensi, err := s.kompetensiRepo.FindByIDs(kompetensiIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to find kompetensi: %w", err)
	}

	if len(kompetensi) != len(kompetensiIDs) {
		return nil, errors.New("one or more kompetensi not found")
	}

	// Convert kompetensi to string representation (comma-separated IDs or names)
	kompetensiStr := ""
	for i, k := range kompetensi {
		if i > 0 {
			kompetensiStr += ","
		}
		kompetensiStr += k.Kode // or use k.Nama or fmt.Sprintf("%d", k.ID) depending on your needs
	}

	// Create new asesor
	asesor := &models.Asesor{
		NamaLengkap:  namaLengkap,
		NoRegistrasi: noRegistrasi,
		Email:        email,
		NoTelepon:    noTelepon,
		Kompetensi:   kompetensiStr,
	}

	err = s.asesorRepo.Create(asesor)
	if err != nil {
		return nil, fmt.Errorf("failed to create asesor: %w", err)
	}

	return asesor, nil
}

func (s *asesorService) UpdateAsesor(id uint, namaLengkap, noRegistrasi, email, noTelepon string, kompetensiIDs []uint) (*models.Asesor, error) {
	// Check if asesor exists
	asesor, err := s.asesorRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("asesor not found: %w", err)
	}

	// Check if registration number is already used by another asesor
	if asesor.NoRegistrasi != noRegistrasi {
		existingAsesor, err := s.asesorRepo.FindByNoRegistrasi(noRegistrasi)
		if err == nil && existingAsesor.ID != id {
			return nil, errors.New("registration number already used by another asesor")
		} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	// Get kompetensi by IDs
	kompetensi, err := s.kompetensiRepo.FindByIDs(kompetensiIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to find kompetensi: %w", err)
	}

	if len(kompetensi) != len(kompetensiIDs) {
		return nil, errors.New("one or more kompetensi not found")
	}

	// Convert kompetensi to string representation
	kompetensiStr := ""
	for i, k := range kompetensi {
		if i > 0 {
			kompetensiStr += ","
		}
		kompetensiStr += k.Kode // or use k.Nama or fmt.Sprintf("%d", k.ID) depending on your needs
	}

	// Update asesor
	asesor.NamaLengkap = namaLengkap
	asesor.NoRegistrasi = noRegistrasi
	asesor.Email = email
	asesor.NoTelepon = noTelepon
	asesor.Kompetensi = kompetensiStr

	err = s.asesorRepo.Update(asesor)
	if err != nil {
		return nil, fmt.Errorf("failed to update asesor: %w", err)
	}

	return asesor, nil
}

func (s *asesorService) DeleteAsesor(id uint) error {
	// Check if asesor exists
	_, err := s.asesorRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("asesor not found: %w", err)
	}

	return s.asesorRepo.Delete(id)
}

func (s *asesorService) GetAsesorByID(id uint) (*models.Asesor, error) {
	return s.asesorRepo.FindByID(id)
}

func (s *asesorService) GetAllAsesors() ([]models.Asesor, error) {
	return s.asesorRepo.FindAll()
}

func (s *asesorService) GetAsesorByNoRegistrasi(noRegistrasi string) (*models.Asesor, error) {
	return s.asesorRepo.FindByNoRegistrasi(noRegistrasi)
}
