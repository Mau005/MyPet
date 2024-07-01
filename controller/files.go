package controller

import (
	"github.com/Mau005/MyPet/db"
	"github.com/Mau005/MyPet/models"
)

type ControllerFiles struct{}

func (cf *ControllerFiles) CreateFiles(files models.Files) (models.Files, error) {
	if err := db.DB.Create(&files).Error; err != nil {
		return files, err
	}
	return files, nil
}

func (cf *ControllerFiles) SaveFiles(files models.Files) (models.Files, error) {
	if err := db.DB.Save(&files).Error; err != nil {
		return files, err
	}
	return files, nil
}

func (cf *ControllerFiles) GetFilesID(idFiles uint) (files models.Files, err error) {
	if err = db.DB.Where("id = ?", idFiles).First(&files).Error; err != nil {
		return
	}
	return
}
