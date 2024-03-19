package repository

import (
	"tour-service/model"

	"gorm.io/gorm"
)

type BundleRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *BundleRepository) Get(id string) (model.Bundle, error) {
	bundle := model.Bundle{}
	dbResult := repo.DatabaseConnection.First(&bundle, "id = ?", id)

	if dbResult.Error != nil {
		return bundle, dbResult.Error
	}

	return bundle, nil
}

func (repo *BundleRepository) GetAll() ([]model.Bundle, error) {
	var bundles []model.Bundle
	dbResult := repo.DatabaseConnection.Find(&bundles)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return bundles, nil
}

func (repo *BundleRepository) Save(bundle *model.Bundle) error {
	dbResult := repo.DatabaseConnection.Create(bundle)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

func (repo *BundleRepository) Update(bundle *model.Bundle) error {
	dbResult := repo.DatabaseConnection.Save(bundle)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *BundleRepository) Delete(id string) error {
	dbResult := repo.DatabaseConnection.Delete(&model.Bundle{}, id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
