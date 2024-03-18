package model

import (
	"github.com/google/uuid"
)

type Status string

const (
	Draft     Status = "draft"
	Published Status = "published"
	Archived  Status = "archived"
)

type Bundle struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"userId"`
	Name   string    `json:"name"`
	Price  float64   `json:"price"`
	Tours  []int     `json:"tours"`
	Status Status    `json:"status"`
	Image  string    `json:"image"`
}
