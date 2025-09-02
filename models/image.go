package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Path   string `json:"path"`
	URL    string `gorm:"-" json:"url"`
}
