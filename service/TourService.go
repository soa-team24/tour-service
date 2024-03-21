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
		Name:        dto.Name,
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

func (service *TourService) Update(tourDto *dto.TourDto) (*dto.TourDto, error) {
	existingTour, err := service.TourRepo.Get(tourDto.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to find blog with ID %s: %v", tourDto.Id, err)
	}

	difficulty_int, err := strconv.Atoi(tourDto.Difficulty)
	if err != nil {
		difficulty_int = 0
	}

	existingTour.Name = tourDto.Name
	existingTour.Description = tourDto.Description
	existingTour.PublishTime = tourDto.PublishTime
	existingTour.Status = model.TourStatus(tourDto.Status)
	existingTour.Image = tourDto.Image
	existingTour.Difficulty = difficulty_int
	existingTour.Price = tourDto.Price
	existingTour.BicycleTime = tourDto.BicycleTime
	existingTour.FootTime = tourDto.FootTime
	existingTour.CarTime = tourDto.CarTime
	existingTour.TotalLength = tourDto.TotalLength

	err = service.TourRepo.Update(&existingTour)
	if err != nil {
		return nil, err
	}
	return tourDto, nil
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

func (service *TourService) GetToursByAuthor(authorID uint32) ([]model.Tour, error) {
	tours, err := service.TourRepo.GetToursByAuthor(authorID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tours by author: %v", err)
	}
	return tours, nil
}
