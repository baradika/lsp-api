package models

import (
	"time"

	"gorm.io/gorm"
)

type FormAPL01 struct {
	ID                   uint           `gorm:"column:id_apl01;primaryKey" json:"id"`
	IDAsesmen            *uint          `gorm:"column:id_asesmen" json:"id_asesmen,omitempty"`
	NamaLengkap          string         `gorm:"size:100" json:"nama_lengkap,omitempty"`
	NoKTP                string         `gorm:"size:20" json:"no_ktp,omitempty"`
	TempatLahir          string         `gorm:"size:50" json:"tempat_lahir,omitempty"`
	TanggalLahir         *time.Time     `json:"tanggal_lahir,omitempty"`
	JenisKelamin         string         `gorm:"type:enum('Laki-laki','Perempuan')" json:"jenis_kelamin,omitempty"`
	Kebangsaan           string         `gorm:"size:50" json:"kebangsaan,omitempty"`
	Alamat               string         `gorm:"type:text" json:"alamat,omitempty"`
	KodePos              string         `gorm:"size:10" json:"kode_pos,omitempty"`
	NoRumah              string         `gorm:"size:20" json:"no_rumah,omitempty"`
	NoKantor             string         `gorm:"size:20" json:"no_kantor,omitempty"`
	NoHP                 string         `gorm:"size:20" json:"no_hp,omitempty"`
	Email                string         `gorm:"size:100" json:"email,omitempty"`
	KualifikasiPendidikan string         `gorm:"size:100" json:"kualifikasi_pendidikan,omitempty"`
	Institusi            string         `gorm:"size:100" json:"institusi,omitempty"`
	Jabatan              string         `gorm:"size:100" json:"jabatan,omitempty"`
	AlamatKantor         string         `gorm:"type:text" json:"alamat_kantor,omitempty"`
	KodePosKantor        string         `gorm:"size:10" json:"kode_pos_kantor,omitempty"`
	TelpKantor           string         `gorm:"size:20" json:"telp_kantor,omitempty"`
	FaxKantor            string         `gorm:"size:20" json:"fax_kantor,omitempty"`
	EmailKantor          string         `gorm:"size:100" json:"email_kantor,omitempty"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"-"`
	Asesmen              *Asesmen       `gorm:"foreignKey:IDAsesmen" json:"asesmen,omitempty"`
}

func (FormAPL01) TableName() string {
	return "form_apl01"
}