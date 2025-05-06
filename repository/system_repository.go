package repository

import (
	"github.com/naufal/simba-qr-app/config"
	"github.com/naufal/simba-qr-app/models"
)

func GetSystemByID(id string) (models.System, error) {
	var system models.System
	err := config.DB.First(&system, "id = ?", id).Error

	return system, err
}

func CreateSystem(system models.System) (models.System, error) {
	err := config.DB.Create(&system).Error
	return system, err
}

func DeleteSystem(id string) (bool error) {
	return config.DB.Delete(&models.System{}, "id = ?", id).Error
}
