package models

import (
	"time"

	"gorm.io/gorm"
)

type Asesor struct {
	ID           uint           `gorm:"column:id_asesor;primaryKey" json:"id"`
	UserID       uint           `gorm:"column:id_user;not null" json:"user_id"`
	NamaLengkap  string         `gorm:"size:100;not null" json:"nama_lengkap"`
	NoRegistrasi string         `gorm:"size:50;uniqueIndex" json:"no_registrasi"`
	Email        string         `gorm:"size:100;uniqueIndex" json:"email"`
	NoTelepon    string         `gorm:"size:20" json:"no_telepon"`
	Kompetensi   string         `gorm:"type:text" json:"kompetensi"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	User         User           `gorm:"foreignKey:UserID" json:"-"`
}

func (Asesor) TableName() string {
	return "asesor"
}
