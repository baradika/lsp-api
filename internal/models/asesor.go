package models

import (
	"time"

	"gorm.io/gorm"
)

type Asesor struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	NamaLengkap  string         `gorm:"size:150;not null" json:"nama_lengkap"`
	NoRegistrasi string         `gorm:"size:50;uniqueIndex;not null" json:"no_registrasi"`
	Email        string         `gorm:"size:100;uniqueIndex;not null" json:"email"`
	NoTelepon    string         `gorm:"size:20" json:"no_telepon"`
	Kompetensi   []Kompetensi   `gorm:"many2many:asesor_kompetensi;" json:"kompetensi"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

type Kompetensi struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Nama        string         `gorm:"size:150;not null" json:"nama"`
	Kode        string         `gorm:"size:50;uniqueIndex;not null" json:"kode"`
	Deskripsi   string         `gorm:"type:text" json:"deskripsi"`
	Asesor      []Asesor       `gorm:"many2many:asesor_kompetensi;" json:"-"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}