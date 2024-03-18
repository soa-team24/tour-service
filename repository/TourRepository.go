package repository

import (
	"errors"
	"tour-service/model"

	"gorm.io/gorm"
)

type TourRepository struct {
	DatabaseConnection *gorm.DB
}

func CheckDBConnection(db *gorm.DB) error {
	if db == nil {
		return errors.New("database connection is nil")
	}
	return nil
}

func (repo *TourRepository) Get(id string) (model.Tour, error) {
	blog := model.Tour{}
	dbResult := repo.DatabaseConnection.First(&blog, "id = ?", id)

	if dbResult.Error != nil {
		return blog, dbResult.Error
	}

	return blog, nil
}

func (repo *TourRepository) GetAll() ([]model.Tour, error) {
	var tours []model.Tour
	dbResult := repo.DatabaseConnection.Find(&tours)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return tours, nil
}

func (repo *TourRepository) Save(tour *model.Tour) error {
	dbResult := repo.DatabaseConnection.Create(tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

func (repo *TourRepository) Update(tour *model.Tour) error {
	dbResult := repo.DatabaseConnection.Save(tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *TourRepository) Delete(id string) error {
	dbResult := repo.DatabaseConnection.Delete(&model.Tour{}, id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
