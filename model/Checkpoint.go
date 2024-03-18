package model

import "github.com/google/uuid"

type Checkpoint struct {
	ID          uuid.UUID `json:"id"`
	TourID      uuid.UUID `json:"tourId"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	IsPublic    bool      `json:"isPublic"`
}
