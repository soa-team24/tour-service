package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BundleStatus string

const (
	Drafted BundleStatus = "draft"
	Publish BundleStatus = "published"
	Archive BundleStatus = "archived"
)

type Bundle struct {
	ID     uuid.UUID    `json:"id"`
	UserID uint32       `json:"userId"`
	Name   string       `json:"name"`
	Price  float64      `json:"price"`
	Status BundleStatus `json:"status"`
	Image  string       `json:"image"`
}

func (bundle *Bundle) BeforeCreate(scope *gorm.DB) error {
	bundle.ID = uuid.New()
	return nil
}
