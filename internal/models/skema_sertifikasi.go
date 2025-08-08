package models

import (
	"time"
)

type SkemaSertifikasi struct {
	ID            uint      `gorm:"column:id_skema;primaryKey" json:"id"`
	JudulSkema    string    `gorm:"size:255;not null" json:"judul_skema"`
	NomorSkema    string    `gorm:"size:100;not null;unique" json:"nomor_skema"`
	JenisSkema    string    `gorm:"type:enum('KKNI','Okupasi','Klaster')" json:"jenis_skema"`
	Deskripsi     string    `gorm:"type:text" json:"deskripsi"`
	TanggalBerlaku *time.Time `json:"tanggal_berlaku"`
	Status        string    `gorm:"type:enum('Aktif','Tidak Aktif');default:'Aktif'" json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (SkemaSertifikasi) TableName() string {
	return "skema_sertifikasi"
}