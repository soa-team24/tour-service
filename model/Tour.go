package model

import (
	"time"

	"github.com/google/uuid"
)

type TourStatus string

const (
	Draft     TourStatus = "draft"
	Published TourStatus = "published"
	Archived  TourStatus = "archived"
)

type Tour struct {
	ID          uuid.UUID  `json:"id"`
	AuthorID    uuid.UUID  `json:"authorId"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	PublishTime time.Time  `json:"publishTime"`
	Status      TourStatus `json:"status"`
	Image       string     `json:"image"`
	Difficulty  int        `json:"difficulty"`
	Price       float64    `json:"price"`
	Tags        []string   `json:"tags"`
	Equipment   []int64    `json:"equipment"`
	CheckPoints []int64    `json:"checkPoints"`
	Objects     []int64    `json:"objects"`
	FootTime    float64    `json:"footTime"`
	BicycleTime float64    `json:"bicycleTime"`
	CarTime     float64    `json:"carTime"`
	TotalLength float64    `json:"totalLenght"`
}
