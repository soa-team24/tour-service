package service

import (
	"fmt"
	"tour-service/model"
	"tour-service/repository"
)

type TourReviewService struct {
	TourReviewRepo *repository.TourReviewRepository
}

func (service *TourReviewService) Get(id string) (*model.TourReview, error) {
	tourReview, err := service.TourReviewRepo.Get(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("tourReview with id %s not found", id))
	}
	return &tourReview, nil
}

func (service *TourReviewService) GetAll() ([]model.TourReview, error) {
	tourReviews, err := service.TourReviewRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tourReviews: %v", err)
	}
	return tourReviews, nil
}

func (service *TourReviewService) Save(tourReview *model.TourReview) error {
	err := service.TourReviewRepo.Save(tourReview)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourReviewService) Update(tourReview *model.TourReview) error {
	existingTourReview, err := service.TourReviewRepo.Get(tourReview.ID.String())
	if err != nil {
		return fmt.Errorf("failed to find tourReview with ID %s: %v", tourReview.ID, err)
	}

	existingTourReview.Comment = tourReview.Comment
	existingTourReview.Grade = tourReview.Grade
	existingTourReview.Images = tourReview.Images
	existingTourReview.ReviewDate = tourReview.ReviewDate

	err = service.TourReviewRepo.Update(&existingTourReview)
	if err != nil {
		return fmt.Errorf("failed to update tour: %v", err)
	}
	return nil
}

func (service *TourReviewService) Delete(id string) error {
	_, err := service.TourReviewRepo.Get(id)
	if err != nil {
		return fmt.Errorf("failed to find tourReview with ID %s: %v", id, err)
	}

	err = service.TourReviewRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete tourReview: %v", err)
	}
	return nil
}
