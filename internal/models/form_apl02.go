package models

import (
	"time"

	"gorm.io/gorm"
)

type FormAPL02 struct {
	ID        uint           `gorm:"column:id_apl02;primaryKey" json:"id"`
	IDAsesmen *uint          `gorm:"column:id_asesmen" json:"id_asesmen,omitempty"`
	KodeUnit  string         `gorm:"size:50" json:"kode_unit,omitempty"`
	JudulUnit string         `gorm:"size:255" json:"judul_unit,omitempty"`
	Elemen    string         `gorm:"type:text" json:"elemen,omitempty"`
	KUK       string         `gorm:"column:kuk;type:text" json:"kuk,omitempty"`
	Status    string         `gorm:"type:enum('K','BK')" json:"status,omitempty"`
	Bukti     string         `gorm:"type:text" json:"bukti,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Asesmen   *Asesmen       `gorm:"foreignKey:IDAsesmen" json:"asesmen,omitempty"`
}

func (FormAPL02) TableName() string {
	return "form_apl02"
}