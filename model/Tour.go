package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TourStatus string

const (
	Draft     TourStatus = "draft"
	Published TourStatus = "published"
	Archived  TourStatus = "archived"
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

func (bundle *Tour) BeforeCreate(scope *gorm.DB) error {
	bundle.ID = uuid.New()
	return nil
}

// SerializeTags serializes the tags field into a JSON string
func (t *Tour) SerializeTags() (string, error) {
	serializedTags, err := json.Marshal(t.Tags)
	if err != nil {
		return "", err
	}
	return string(serializedTags), nil
}

// DeserializeTags deserializes the JSON string into a slice of strings
func (t *Tour) DeserializeTags(serializedTags string) error {
	return json.Unmarshal([]byte(serializedTags), &t.Tags)
}
