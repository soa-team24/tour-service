package service

import (
	"fmt"
	"tour-service/model"
	"tour-service/repository"
)

type EquipmentService struct {
	EquipmentRepo *repository.EquipmentRepository
}

func (service *EquipmentService) Get(id string) (*model.Equipment, error) {
	equipment, err := service.EquipmentRepo.Get(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("equipment with id %s not found", id))
	}
	return &equipment, nil
}

func (service *EquipmentService) GetAll() ([]model.Equipment, error) {
	equipments, err := service.EquipmentRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve equipments: %v", err)
	}
	return equipments, nil
}

func (service *EquipmentService) Save(equipment *model.Equipment) error {
	err := service.EquipmentRepo.Save(equipment)
	if err != nil {
		return err
	}
	return nil
}

func (service *EquipmentService) Update(equipment *model.Equipment) error {
	existingEquipment, err := service.EquipmentRepo.Get(fmt.Sprintf("%d", equipment.ID)) //proveri
	if err != nil {
		return fmt.Errorf("failed to find equipment with ID %d: %v", equipment.ID, err)
	}

	existingEquipment.ID = equipment.ID
	existingEquipment.Name = equipment.Name
	existingEquipment.Description = equipment.Description

	err = service.EquipmentRepo.Update(&existingEquipment)
	if err != nil {
		return fmt.Errorf("failed to update equipment: %v", err)
	}
	return nil
}

func (service *EquipmentService) Delete(id string) error {
	_, err := service.EquipmentRepo.Get(id)
	if err != nil {
		return fmt.Errorf("failed to find equipment with ID %s: %v", id, err)
	}

	err = service.EquipmentRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete equipment: %v", err)
	}
	return nil
}
