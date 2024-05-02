package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name          string     `gorm:"unique" json:"name"`
	Password      string     `gorm:"not null" json:"-"`
	Email         string     `gorm:"unique" json:"email"`
	ActiveAccount bool       `gorm:"default=0" json:"activeAccount"`
	Language      string     `json:"languaje"`
	Access        uint8      `gorm:"not null, default=0" json:"access"`
	Premmiun      *time.Time `json:"premmiun"`
}
