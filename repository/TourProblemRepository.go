package repository

import (
	"tour-service/model"

	"gorm.io/gorm"
)

type TourProblemRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourProblemRepository) Get(id string) (model.TourProblem, error) {
	tourProblem := model.TourProblem{}
	dbResult := repo.DatabaseConnection.First(&tourProblem, "id = ?", id)

	if dbResult.Error != nil {
		return tourProblem, dbResult.Error
	}

	return tourProblem, nil
}

func (repo *TourProblemRepository) GetAll() ([]model.TourProblem, error) {
	var tourProblems []model.TourProblem
	dbResult := repo.DatabaseConnection.Find(&tourProblems)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return tourProblems, nil
}

func (repo *TourProblemRepository) Save(tourProblem *model.TourProblem) error {
	dbResult := repo.DatabaseConnection.Create(tourProblem)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

func (repo *TourProblemRepository) Update(tourProblem *model.TourProblem) error {
	dbResult := repo.DatabaseConnection.Save(tourProblem)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *TourProblemRepository) Delete(id string) error {
	dbResult := repo.DatabaseConnection.Delete(&model.TourProblem{}, "id = ?", id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *TourProblemRepository) GetTourProblemsForTourist(touristIdId uint32) ([]model.TourProblem, error) {
	var tourProblems []model.TourProblem
	result := repo.DatabaseConnection.Find(&tourProblems, "tourist_id = ?", touristIdId)
	if result.Error != nil {
		return nil, result.Error
	}
	return tourProblems, nil
}
