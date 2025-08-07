package models

import (
	"time"
)

type Jurusan struct {
	ID          uint      `gorm:"column:id_jurusan;primaryKey" json:"id"`
	KodeJurusan string    `gorm:"size:20;not null;unique" json:"kode_jurusan"`
	NamaJurusan string    `gorm:"size:100;not null" json:"nama_jurusan"`
	Jenjang     string    `gorm:"type:enum('D1','D2','D3','D4','S1','S2','S3','SMK','SMA');not null" json:"jenjang"`
	Deskripsi   string    `gorm:"type:text" json:"deskripsi"`
	Status      string    `gorm:"type:enum('Aktif','Non-Aktif');default:'Aktif'" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Jurusan) TableName() string {
	return "jurusan"
}
