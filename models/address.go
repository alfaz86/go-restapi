package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}
