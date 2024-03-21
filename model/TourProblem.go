package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TourProblem struct {
	ID                uuid.UUID  `json:"id"`
	ProblemCategory   string     `json:"problemCategory"`
	ProblemPriority   string     `json:"problemPriority"`
	Description       string     `json:"description"`
	TimeStamp         time.Time  `json:"timeStamp"`
	TourId            string     `json:"tourId"`
	IsClosed          bool       `json:"isClosed"`
	IsResolved        bool       `json:"isResolved"`
	TouristId         uint32     `json:"touristId"`
	DeadlineTimeStamp *time.Time `json:"deadlineTimeStamp,omitempty"`
}

func (tourProblem *TourProblem) BeforeCreate(scope *gorm.DB) error {
	tourProblem.ID = uuid.New()
	return nil
}
