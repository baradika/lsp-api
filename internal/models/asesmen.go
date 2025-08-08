package models

import (
	"time"

	"gorm.io/gorm"
)

type Asesmen struct {
	ID             uint           `gorm:"column:id_asesmen;primaryKey" json:"id"`
	IDSkema        *uint          `gorm:"column:id_skema" json:"id_skema,omitempty"`
	IDAsesi        *uint          `gorm:"column:id_asesi" json:"id_asesi,omitempty"`
	IDAsesor       *uint          `gorm:"column:id_asesor" json:"id_asesor,omitempty"`
	TanggalMulai   *time.Time     `json:"tanggal_mulai,omitempty"`
	TanggalSelesai *time.Time     `json:"tanggal_selesai,omitempty"`
	TUK            string         `gorm:"column:tuk;size:255" json:"tuk,omitempty"`
	Status         string         `gorm:"type:enum('Draft','Berjalan','Selesai','Ditolak');default:'Draft'" json:"status"`
	Hasil          string         `gorm:"type:enum('Kompeten','Belum Kompeten','Belum Selesai');default:'Belum Selesai'" json:"hasil"`
	Catatan        string         `gorm:"type:text" json:"catatan,omitempty"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	Asesi          *Asesi         `gorm:"foreignKey:IDAsesi" json:"asesi,omitempty"`
	Asesor         *Asesor        `gorm:"foreignKey:IDAsesor" json:"asesor,omitempty"`
}

func (Asesmen) TableName() string {
	return "asesmen"
}