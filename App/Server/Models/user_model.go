package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Cc       int    `gorm:"not null"`
	Name     string `gorm:"not null"`
	Address  string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Archives []Archive
}
