package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TourStatus uint32

const (
	Draft     TourStatus = 0
	Published TourStatus = 1
	Archived  TourStatus = 2
)

type Tour struct {
	ID          uuid.UUID  `json:"id"`
	AuthorID    uint32     `json:"authorId"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	PublishTime string     `json:"publishTime"`
	Status      TourStatus `json:"status"`
	Image       string     `json:"image"`
	Difficulty  int        `json:"difficulty"`
	Price       float64    `json:"price"`
	FootTime    float64    `json:"footTime"`
	BicycleTime float64    `json:"bicycleTime"`
	CarTime     float64    `json:"carTime"`
	TotalLength float64    `json:"totalLength"`
}

func (tour *Tour) BeforeCreate(scope *gorm.DB) error {
	tour.ID = uuid.New()
	return nil
}
