package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BundleTour struct {
	ID       uuid.UUID `json:"id"`
	BundleID uint32    `json:"bundleId"`
	TourID   uint32    `json:"tourId"`
}

func (bundleTour *BundleTour) BeforeCreate(scope *gorm.DB) error {
	bundleTour.ID = uuid.New()
	return nil
}
