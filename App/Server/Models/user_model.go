package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	CivID             int        `gorm:"unique;not null"`
	Name              string     `gorm:"not null"`
	Address           string     `gorm:"not null"`
	Email             string     `gorm:"not null"`
	Password          *string    `json:"password"`
	Token             string     `json:"token,omitempty"`
	PasswordCreatedAt *time.Time `json:"-"`
	Archives          []Archive  `gorm:"foreignKey:Owner;references:CivID"`
}

type Citizen struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	OperatorId   int    `json:"operatorId"`
	OperatorName string `json:"operatorName"`
}

type GenerativeInfo struct {
	CivID int    `json:"id"`
	Email string `json:"email"`
}

type LoginInfo struct {
	CivID    int    `json:"id"`
	Password string `json:"password"`
}
