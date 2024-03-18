package repository

import (
	"tour-service/model"

	"gorm.io/gorm"
)

type EquipmentRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *EquipmentRepository) Get(id string) (model.Equipment, error) {
	blog := model.Equipment{}
	dbResult := repo.DatabaseConnection.First(&blog, "id = ?", id)

	if dbResult.Error != nil {
		return blog, dbResult.Error
	}

	return blog, nil
}

func (repo *EquipmentRepository) GetAll() ([]model.Equipment, error) {
	var equipment []model.Equipment
	dbResult := repo.DatabaseConnection.Find(&equipment)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return equipment, nil
}

func (repo *EquipmentRepository) Save(equipment *model.Equipment) error {
	dbResult := repo.DatabaseConnection.Create(equipment)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

func (repo *EquipmentRepository) Update(equipment *model.Equipment) error {
	dbResult := repo.DatabaseConnection.Save(equipment)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *EquipmentRepository) Delete(id string) error {
	dbResult := repo.DatabaseConnection.Delete(&model.Equipment{}, id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
