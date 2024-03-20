package repository

import (
	"tour-service/model"

	"gorm.io/gorm"
)

type TourReviewRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourReviewRepository) Get(id string) (model.TourReview, error) {
	tourReview := model.TourReview{}
	dbResult := repo.DatabaseConnection.First(&tourReview, "id = ?", id)

	if dbResult.Error != nil {
		return tourReview, dbResult.Error
	}

	return tourReview, nil
}

func (repo *TourReviewRepository) GetAll() ([]model.TourReview, error) {
	var tourReviews []model.TourReview
	dbResult := repo.DatabaseConnection.Find(&tourReviews)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return tourReviews, nil
}

func (repo *TourReviewRepository) Save(tourReview *model.TourReview) error {
	dbResult := repo.DatabaseConnection.Create(tourReview)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

func (repo *TourReviewRepository) Update(tourReview *model.TourReview) error {
	dbResult := repo.DatabaseConnection.Save(tourReview)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *TourReviewRepository) Delete(id string) error {
	dbResult := repo.DatabaseConnection.Delete(&model.TourReview{}, "id = ?", id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *TourReviewRepository) GetTourReviewsByTourID(tourId string) ([]model.TourReview, error) {
	var tourReviews []model.TourReview
	result := repo.DatabaseConnection.Find(&tourReviews, "tour_id = ?", tourId)
	if result.Error != nil {
		return nil, result.Error
	}
	return tourReviews, nil
}
