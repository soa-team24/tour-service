package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Checkpoint struct {
	ID          uuid.UUID `json:"id"`
	TourID      uint32    `json:"tourId"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	IsPublic    bool      `json:"isPublic"`
}

func (checkpoint *Checkpoint) BeforeCreate(scope *gorm.DB) error {
	checkpoint.ID = uuid.New()
	return nil
}
