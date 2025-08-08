package models

import (
	"time"
)

type KriteriaUnjukKerja struct {
	ID              uint              `gorm:"column:id_kuk;primaryKey" json:"id"`
	IDElemen        *uint             `gorm:"column:id_elemen" json:"id_elemen,omitempty"`
	DeskripsiKUK    string            `gorm:"type:text;not null" json:"deskripsi_kuk"`
	Urutan          *int              `json:"urutan,omitempty"`
	StandarIndustri string            `gorm:"type:text" json:"standar_industri"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	Elemen          *ElemenKompetensi `gorm:"foreignKey:IDElemen" json:"elemen,omitempty"`
}

func (KriteriaUnjukKerja) TableName() string {
	return "kriteria_unjuk_kerja"
}