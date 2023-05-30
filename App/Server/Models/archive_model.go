package models

import "gorm.io/gorm"

type Archive struct {
	gorm.Model

	UserID uint   `json:"userId" xml:"userId"`
	Name   string `json:"name" xml:"name"`
	Type   string `json:"type" xml:"type"`
	Path   string `json:"path" xml:"path"`
}
