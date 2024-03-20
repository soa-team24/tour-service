package service

import (
	"fmt"
	"strconv"
	"tour-service/dto"
	"tour-service/model"
	"tour-service/repository"
)

type TourService struct {
	TourRepo *repository.TourRepository
}

func (service *TourService) Get(id string) (*model.Tour, error) {
	tour, err := service.TourRepo.Get(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("tour with id %s not found", id))
	}
	return &tour, nil
}

func (service *TourService) GetAll() ([]model.Tour, error) {
	tours, err := service.TourRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tours: %v", err)
	}
	return tours, nil
}

func (service *TourService) Save(dto *dto.TourDto) (*model.Tour, error) {
	difficultyInt, err := strconv.Atoi(dto.Difficulty)
	if err != nil {
		difficultyInt = 0
	}
	tour := &model.Tour{
		Title:       dto.Name,
		Description: dto.Description,
		PublishTime: dto.PublishTime,
		Status:      model.TourStatus(dto.Status),
		Image:       dto.Image,
		Difficulty:  difficultyInt,
		Price:       dto.Price,
		FootTime:    dto.FootTime,
		BicycleTime: dto.BicycleTime,
		CarTime:     dto.CarTime,
		TotalLength: dto.TotalLength,
		AuthorID:    dto.AuthorID,
	}

	checkpoint, err := service.TourRepo.Save(tour)
	if err != nil {
		return nil, err
	}
	return checkpoint, nil
}

func (service *TourService) Update(tour *model.Tour) error {
	existingTour, err := service.TourRepo.Get(tour.ID.String())
	if err != nil {
		return fmt.Errorf("failed to find blog with ID %s: %v", tour.ID, err)
	}

	existingTour.Title = tour.Title
	existingTour.Description = tour.Description
	existingTour.PublishTime = tour.PublishTime
	existingTour.Status = tour.Status
	existingTour.Image = tour.Image
	existingTour.Difficulty = tour.Difficulty
	existingTour.Price = tour.Price
	existingTour.BicycleTime = tour.BicycleTime
	existingTour.FootTime = tour.FootTime
	existingTour.CarTime = tour.CarTime
	existingTour.TotalLength = tour.TotalLength

	err = service.TourRepo.Update(&existingTour)
	if err != nil {
		return fmt.Errorf("failed to update tour: %v", err)
	}
	return nil
}

func (service *TourService) Delete(id string) error {
	_, err := service.TourRepo.Get(id)
	if err != nil {
		return fmt.Errorf("failed to find tour with ID %s: %v", id, err)
	}

	err = service.TourRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete tour: %v", err)
	}
	return nil
}
