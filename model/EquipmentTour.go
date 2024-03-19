package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EquipmentTour struct {
	ID          uuid.UUID `json:"id"`
	EquipmentID uint32    `json:"bundleId"`
	TourID      uint32    `json:"tourId"`
}

func (equipmentTour *EquipmentTour) BeforeCreate(scope *gorm.DB) error {
	equipmentTour.ID = uuid.New()
	return nil
}
