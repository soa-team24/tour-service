package repository

import (
	"tour-service/model"

	"gorm.io/gorm"
)

type CheckpointRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *CheckpointRepository) Get(id string) (model.Checkpoint, error) {
	checkpoint := model.Checkpoint{}
	dbResult := repo.DatabaseConnection.First(&checkpoint, "id = ?", id)

	if dbResult.Error != nil {
		return checkpoint, dbResult.Error
	}

	return checkpoint, nil
}

func (repo *CheckpointRepository) GetAll() ([]model.Checkpoint, error) {
	var checkpoints []model.Checkpoint
	dbResult := repo.DatabaseConnection.Find(&checkpoints)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return checkpoints, nil
}

func (repo *CheckpointRepository) Save(checkpoint *model.Checkpoint) (*model.Checkpoint, error) {
	dbResult := repo.DatabaseConnection.Create(checkpoint)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return checkpoint, nil
}

func (repo *CheckpointRepository) Update(checkpoint *model.Checkpoint) error {
	dbResult := repo.DatabaseConnection.Save(checkpoint)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *CheckpointRepository) Delete(id string) error {
	dbResult := repo.DatabaseConnection.Delete(&model.Checkpoint{}, "id = ?", id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
