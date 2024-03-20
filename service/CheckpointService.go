package service

import (
	"fmt"
	"tour-service/model"
	"tour-service/repository"
)

type CheckpointService struct {
	CheckpointRepo *repository.CheckpointRepository
}

func (service *CheckpointService) Get(id string) (*model.Checkpoint, error) {
	checkpoint, err := service.CheckpointRepo.Get(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("checkpoint with id %s not found", id))
	}
	return &checkpoint, nil
}

func (service *CheckpointService) GetAll() ([]model.Checkpoint, error) {
	checkpoints, err := service.CheckpointRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve checkpoints: %v", err)
	}
	return checkpoints, nil
}

func (service *CheckpointService) Save(checkpoint *model.Checkpoint) (*model.Checkpoint, error) {
	checkpoint, err := service.CheckpointRepo.Save(checkpoint)
	if err != nil {
		return nil, err
	}
	return checkpoint, nil
}

func (service *CheckpointService) Update(checkpoint *model.Checkpoint) error {
	existingCheckpoint, err := service.CheckpointRepo.Get(checkpoint.ID.String())
	if err != nil {
		return fmt.Errorf("failed to find checkpoint with ID %s: %v", checkpoint.ID, err)
	}

	existingCheckpoint.ID = checkpoint.ID
	existingCheckpoint.TourID = checkpoint.TourID
	existingCheckpoint.Latitude = checkpoint.Latitude
	existingCheckpoint.Name = checkpoint.Name
	existingCheckpoint.Description = checkpoint.Description
	existingCheckpoint.Image = checkpoint.Image
	existingCheckpoint.IsPublic = checkpoint.IsPublic

	err = service.CheckpointRepo.Update(&existingCheckpoint)
	if err != nil {
		return fmt.Errorf("failed to update checkpoint: %v", err)
	}
	return nil
}

func (service *CheckpointService) Delete(id string) error {
	_, err := service.CheckpointRepo.Get(id)
	if err != nil {
		return fmt.Errorf("failed to find checkpoint with ID %s: %v", id, err)
	}

	err = service.CheckpointRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete checkpoint: %v", err)
	}
	return nil
}

func (service *CheckpointService) GetCheckpointsByTourID(tourId string) ([]model.Checkpoint, error) {
	checkpoints, err := service.CheckpointRepo.GetCheckpointsByTourID(tourId)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve checkpoints for tour with ID %s: %v", tourId, err)
	}
	return checkpoints, nil
}
