package models

import "gorm.io/gorm"

type Archive struct {
	gorm.Model

	Owner           int    `gorm:"not null"`
	FullName        string `gorm:"not null"`
	Name            string `gorm:"not null"`
	Type            string `gorm:"not null"`
	Path            string `gorm:"not null"`
	IsAuthenticated bool   `gorm:"not null"`
}
