package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                uint       `gorm:"column:id_user;primaryKey" json:"id"`
	Username          string     `gorm:"size:50;not null;unique" json:"username"`
	Password          string     `gorm:"column:password_hash;size:255;not null" json:"-"`
	Email             string     `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Role              string     `gorm:"type:enum('Admin','Asesor','Asesi');not null" json:"role"`
	IDRelated         *uint      `gorm:"column:id_related" json:"id_related,omitempty"`
	LastLogin         *time.Time `json:"last_login,omitempty"`
	IsActive          bool       `gorm:"default:1" json:"is_active"`
	ResetToken        *string    `gorm:"size:100" json:"-"`
	ResetTokenExpires *time.Time `json:"-"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

// ComparePassword compares the provided password with the hashed password
func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// HashPassword hashes the provided password
func (u *User) HashPassword(password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}
