package models

import "gorm.io/gorm"

type Archive struct {
	gorm.Model

	CivID int    `gorm:"not null"`
	Name  string `gorm:"not null"`
	Type  string `gorm:"not null"`
	Path  string `gorm:"not null"`
}
