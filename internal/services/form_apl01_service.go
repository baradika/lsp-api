package services

import (
	"errors"
	"fmt"
	"time"

	"lsp-api/internal/models"
	"lsp-api/internal/repositories"

	"gorm.io/gorm"
)

type FormAPL01Service interface {
	CreateFormAPL01(idAsesmen uint, data map[string]interface{}) (*models.FormAPL01, error)
	UpdateFormAPL01(id uint, data map[string]interface{}) (*models.FormAPL01, error)
	DeleteFormAPL01(id uint) error
	GetFormAPL01ByID(id uint) (*models.FormAPL01, error)
	GetFormAPL01ByAsesmenID(asesmenID uint) (*models.FormAPL01, error)
	GetFormAPL01ByAsesiID(asesiID uint) ([]models.FormAPL01, error)
}

type formAPL01Service struct {
	formRepo    repositories.FormAPL01Repository
	asesmenRepo repositories.AsesmenRepository
}

func NewFormAPL01Service(
	formRepo repositories.FormAPL01Repository,
	asesmenRepo repositories.AsesmenRepository,
) FormAPL01Service {
	return &formAPL01Service{
		formRepo:    formRepo,
		asesmenRepo: asesmenRepo,
	}
}

func (s *formAPL01Service) CreateFormAPL01(idAsesmen uint, data map[string]interface{}) (*models.FormAPL01, error) {
	// Verify asesmen exists
	_, err := s.asesmenRepo.FindByID(idAsesmen)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("asesmen not found")
		}
		return nil, err
	}

	// Check if form already exists for this asesmen
	existingForm, err := s.formRepo.FindByAsesmenID(idAsesmen)
	if err == nil && existingForm != nil {
		return nil, errors.New("form APL01 already exists for this asesmen")
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Create new form
	form := &models.FormAPL01{
		IDAsesmen: &idAsesmen,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Set fields from data map
	if namaLengkap, ok := data["nama_lengkap"].(string); ok {
		form.NamaLengkap = namaLengkap
	}
	if noKTP, ok := data["no_ktp"].(string); ok {
		form.NoKTP = noKTP
	}
	if tempatLahir, ok := data["tempat_lahir"].(string); ok {
		form.TempatLahir = tempatLahir
	}
	if tanggalLahirStr, ok := data["tanggal_lahir"].(string); ok {
		tanggalLahir, err := time.Parse("2006-01-02", tanggalLahirStr)
		if err == nil {
			form.TanggalLahir = &tanggalLahir
		}
	}
	if jenisKelamin, ok := data["jenis_kelamin"].(string); ok {
		form.JenisKelamin = jenisKelamin
	}
	if kebangsaan, ok := data["kebangsaan"].(string); ok {
		form.Kebangsaan = kebangsaan
	}
	if alamat, ok := data["alamat"].(string); ok {
		form.Alamat = alamat
	}
	if kodePos, ok := data["kode_pos"].(string); ok {
		form.KodePos = kodePos
	}
	if noRumah, ok := data["no_rumah"].(string); ok {
		form.NoRumah = noRumah
	}
	if noKantor, ok := data["no_kantor"].(string); ok {
		form.NoKantor = noKantor
	}
	if noHP, ok := data["no_hp"].(string); ok {
		form.NoHP = noHP
	}
	if email, ok := data["email"].(string); ok {
		form.Email = email
	}
	if kualifikasiPendidikan, ok := data["kualifikasi_pendidikan"].(string); ok {
		form.KualifikasiPendidikan = kualifikasiPendidikan
	}
	if institusi, ok := data["institusi"].(string); ok {
		form.Institusi = institusi
	}
	if jabatan, ok := data["jabatan"].(string); ok {
		form.Jabatan = jabatan
	}
	if alamatKantor, ok := data["alamat_kantor"].(string); ok {
		form.AlamatKantor = alamatKantor
	}
	if kodePosKantor, ok := data["kode_pos_kantor"].(string); ok {
		form.KodePosKantor = kodePosKantor
	}
	if telpKantor, ok := data["telp_kantor"].(string); ok {
		form.TelpKantor = telpKantor
	}
	if faxKantor, ok := data["fax_kantor"].(string); ok {
		form.FaxKantor = faxKantor
	}
	if emailKantor, ok := data["email_kantor"].(string); ok {
		form.EmailKantor = emailKantor
	}

	err = s.formRepo.Create(form)
	if err != nil {
		return nil, fmt.Errorf("failed to create form APL01: %w", err)
	}

	return form, nil
}

func (s *formAPL01Service) UpdateFormAPL01(id uint, data map[string]interface{}) (*models.FormAPL01, error) {
	// Check if form exists
	form, err := s.formRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("form APL01 not found: %w", err)
	}

	// Update fields from data map
	if namaLengkap, ok := data["nama_lengkap"].(string); ok {
		form.NamaLengkap = namaLengkap
	}
	if noKTP, ok := data["no_ktp"].(string); ok {
		form.NoKTP = noKTP
	}
	if tempatLahir, ok := data["tempat_lahir"].(string); ok {
		form.TempatLahir = tempatLahir
	}
	if tanggalLahirStr, ok := data["tanggal_lahir"].(string); ok {
		tanggalLahir, err := time.Parse("2006-01-02", tanggalLahirStr)
		if err == nil {
			form.TanggalLahir = &tanggalLahir
		}
	}
	if jenisKelamin, ok := data["jenis_kelamin"].(string); ok {
		form.JenisKelamin = jenisKelamin
	}
	if kebangsaan, ok := data["kebangsaan"].(string); ok {
		form.Kebangsaan = kebangsaan
	}
	if alamat, ok := data["alamat"].(string); ok {
		form.Alamat = alamat
	}
	if kodePos, ok := data["kode_pos"].(string); ok {
		form.KodePos = kodePos
	}
	if noRumah, ok := data["no_rumah"].(string); ok {
		form.NoRumah = noRumah
	}
	if noKantor, ok := data["no_kantor"].(string); ok {
		form.NoKantor = noKantor
	}
	if noHP, ok := data["no_hp"].(string); ok {
		form.NoHP = noHP
	}
	if email, ok := data["email"].(string); ok {
		form.Email = email
	}
	if kualifikasiPendidikan, ok := data["kualifikasi_pendidikan"].(string); ok {
		form.KualifikasiPendidikan = kualifikasiPendidikan
	}
	if institusi, ok := data["institusi"].(string); ok {
		form.Institusi = institusi
	}
	if jabatan, ok := data["jabatan"].(string); ok {
		form.Jabatan = jabatan
	}
	if alamatKantor, ok := data["alamat_kantor"].(string); ok {
		form.AlamatKantor = alamatKantor
	}
	if kodePosKantor, ok := data["kode_pos_kantor"].(string); ok {
		form.KodePosKantor = kodePosKantor
	}
	if telpKantor, ok := data["telp_kantor"].(string); ok {
		form.TelpKantor = telpKantor
	}
	if faxKantor, ok := data["fax_kantor"].(string); ok {
		form.FaxKantor = faxKantor
	}
	if emailKantor, ok := data["email_kantor"].(string); ok {
		form.EmailKantor = emailKantor
	}

	form.UpdatedAt = time.Now()

	err = s.formRepo.Update(form)
	if err != nil {
		return nil, fmt.Errorf("failed to update form APL01: %w", err)
	}

	return form, nil
}

func (s *formAPL01Service) DeleteFormAPL01(id uint) error {
	// Check if form exists
	_, err := s.formRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("form APL01 not found: %w", err)
	}

	return s.formRepo.Delete(id)
}

func (s *formAPL01Service) GetFormAPL01ByID(id uint) (*models.FormAPL01, error) {
	return s.formRepo.FindByID(id)
}

func (s *formAPL01Service) GetFormAPL01ByAsesmenID(asesmenID uint) (*models.FormAPL01, error) {
	return s.formRepo.FindByAsesmenID(asesmenID)
}

func (s *formAPL01Service) GetFormAPL01ByAsesiID(asesiID uint) ([]models.FormAPL01, error) {
	return s.formRepo.FindByAsesiID(asesiID)
}