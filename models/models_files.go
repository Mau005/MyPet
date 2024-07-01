package models

import "gorm.io/gorm"

type Files struct {
	gorm.Model
	Name        string
	Description string
	Url_Path    string
}
