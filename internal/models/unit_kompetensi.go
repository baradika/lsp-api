package models

import (
	"time"
)

type UnitKompetensi struct {
	ID         uint              `gorm:"column:id_unit;primaryKey" json:"id"`
	KodeUnit   string            `gorm:"size:50;not null;unique" json:"kode_unit"`
	JudulUnit  string            `gorm:"size:255;not null" json:"judul_unit"`
	IDSkema    *uint             `gorm:"column:id_skema" json:"id_skema,omitempty"`
	Deskripsi  string            `gorm:"type:text" json:"deskripsi"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	Skema      *SkemaSertifikasi `gorm:"foreignKey:IDSkema" json:"skema,omitempty"`
}

func (UnitKompetensi) TableName() string {
	return "unit_kompetensi"
}