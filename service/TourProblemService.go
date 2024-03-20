package service

import (
	"fmt"
	"tour-service/model"
	"tour-service/repository"
)

type TourProblemService struct {
	TourProblemRepo *repository.TourProblemRepository
}

func (service *TourProblemService) Get(id string) (*model.TourProblem, error) {
	tourProblem, err := service.TourProblemRepo.Get(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("tourProblem with id %s not found", id))
	}
	return &tourProblem, nil
}

func (service *TourProblemService) GetAll() ([]model.TourProblem, error) {
	tourProblems, err := service.TourProblemRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tourProblems: %v", err)
	}
	return tourProblems, nil
}

func (service *TourProblemService) Save(tourProblem *model.TourProblem) error {
	err := service.TourProblemRepo.Save(tourProblem)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourProblemService) Update(tourProblem *model.TourProblem) error {
	existingTourProblem, err := service.TourProblemRepo.Get(tourProblem.ID.String())
	if err != nil {
		return fmt.Errorf("failed to find tourProblem with ID %s: %v", tourProblem.ID, err)
	}

	existingTourProblem.ProblemCategory = tourProblem.ProblemCategory
	existingTourProblem.ProblemPriority = tourProblem.ProblemPriority
	existingTourProblem.Description = tourProblem.Description
	existingTourProblem.IsClosed = tourProblem.IsClosed
	existingTourProblem.IsResolved = tourProblem.IsResolved
	existingTourProblem.DeadlineTimeStamp = tourProblem.DeadlineTimeStamp

	err = service.TourProblemRepo.Update(&existingTourProblem)
	if err != nil {
		return fmt.Errorf("failed to update tourProblem: %v", err)
	}
	return nil
}

func (service *TourProblemService) Delete(id string) error {
	_, err := service.TourProblemRepo.Get(id)
	if err != nil {
		return fmt.Errorf("failed to find tourProblem with ID %s: %v", id, err)
	}

	err = service.TourProblemRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete tourProblem: %v", err)
	}
	return nil
}
