package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TourReview struct {
	ID         uuid.UUID `json:"id"`
	Grade      uint32    `json:"grade"`
	Comment    string    `json:"comment"`
	UserID     uint32    `json:"userId"`
	VisitDate  time.Time `json:"visitDate"`
	ReviewDate time.Time `json:"reviewDate"`
	Images     string    `json:"images"`
	TourID     uuid.UUID `json:"tourId"`
}

func (tourReview *TourReview) BeforeCreate(scope *gorm.DB) error {
	tourReview.ID = uuid.New()
	return nil
}
