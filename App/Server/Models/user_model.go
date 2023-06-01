package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	CivID             int        `gorm:"not null"`
	Name              string     `gorm:"not null"`
	Address           string     `gorm:"not null"`
	Email             string     `gorm:"not null"`
	Password          *string    `json:"password"`
	Token             string     `json:"token,omitempty"`
	PasswordCreatedAt *time.Time `json:"-"`
	Archives          []Archive  `gorm:"foreignKey:UserID"`
}
