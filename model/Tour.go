package model

import (
	"time"

	"github.com/google/uuid"
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
	Title       string     `json:"title"`
	Description string     `json:"description"`
	PublishTime time.Time  `json:"publishTime"`
	Status      TourStatus `json:"status"`
	Image       string     `json:"image"`
	Difficulty  int        `json:"difficulty"`
	Price       float64    `json:"price"`
	Tags        []string   `gorm:"type:text" json:"tags"`
	FootTime    float64    `json:"footTime"`
	BicycleTime float64    `json:"bicycleTime"`
	CarTime     float64    `json:"carTime"`
	TotalLength float64    `json:"totalLength"`
}
