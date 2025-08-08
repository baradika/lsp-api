package models

import (
	"time"
)

type ElemenKompetensi struct {
	ID          uint             `gorm:"column:id_elemen;primaryKey" json:"id"`
	IDUnit      *uint            `gorm:"column:id_unit" json:"id_unit,omitempty"`
	NamaElemen  string           `gorm:"size:255;not null" json:"nama_elemen"`
	Urutan      *int             `json:"urutan,omitempty"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	Unit        *UnitKompetensi  `gorm:"foreignKey:IDUnit" json:"unit,omitempty"`
}

func (ElemenKompetensi) TableName() string {
	return "elemen_kompetensi"
}