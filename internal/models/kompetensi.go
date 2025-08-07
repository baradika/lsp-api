package models

import (
	"time"

	"gorm.io/gorm"
)

type Kompetensi struct {
	ID          uint           `gorm:"column:id_kompetensi;primaryKey" json:"id"`
	Kode        string         `gorm:"size:50;not null;unique" json:"kode"`
	Nama        string         `gorm:"size:100;not null" json:"nama"`
	Deskripsi   string         `gorm:"type:text" json:"deskripsi"`
	Jenis       string         `gorm:"size:50" json:"jenis"`
	Level       string         `gorm:"size:20" json:"level"`
	Status      string         `gorm:"type:enum('Aktif','Non-Aktif');default:'Aktif'" json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Kompetensi) TableName() string {
	return "kompetensi"
}