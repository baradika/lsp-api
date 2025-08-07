package models

import (
	"time"

	"gorm.io/gorm"
)

type Asesi struct {
	ID                    uint           `gorm:"column:id_asesi;primaryKey" json:"id"`
	UserID                uint           `gorm:"column:id_user" json:"user_id"`
	NamaLengkap           string         `gorm:"size:100;not null" json:"nama_lengkap"`
	NoKTP                 string         `gorm:"size:20;uniqueIndex" json:"no_ktp"`
	TempatLahir           string         `gorm:"size:50" json:"tempat_lahir"`
	TanggalLahir          *time.Time     `json:"tanggal_lahir"`
	JenisKelamin          string         `gorm:"type:enum('Laki-laki','Perempuan')" json:"jenis_kelamin"`
	Alamat                string         `gorm:"type:text" json:"alamat"`
	KodePos               string         `gorm:"size:10" json:"kode_pos"`
	NoTelepon             string         `gorm:"size:20" json:"no_telepon"`
	Email                 string         `gorm:"size:100;uniqueIndex" json:"email"`
	KualifikasiPendidikan string         `gorm:"size:100" json:"kualifikasi_pendidikan"`
	IDJurusan             *uint          `gorm:"column:id_jurusan" json:"id_jurusan,omitempty"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `gorm:"index" json:"-"`
	User                  User           `gorm:"foreignKey:UserID" json:"-"`
	Jurusan               *Jurusan       `gorm:"foreignKey:IDJurusan" json:"jurusan,omitempty"`
}

func (Asesi) TableName() string {
	return "asesi"
}
