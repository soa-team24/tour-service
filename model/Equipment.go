package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Equipment struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func (equipment *Equipment) BeforeCreate(scope *gorm.DB) error {
	equipment.ID = uuid.New()
	return nil
}
