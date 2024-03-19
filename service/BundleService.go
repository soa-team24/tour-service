package service

import (
	"fmt"
	"tour-service/model"
	"tour-service/repository"
)

type BundleService struct {
	BundleRepo *repository.BundleRepository
}

func (service *BundleService) Get(id string) (*model.Bundle, error) {
	bundle, err := service.BundleRepo.Get(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("bundle with id %s not found", id))
	}
	return &bundle, nil
}

func (service *BundleService) GetAll() ([]model.Bundle, error) {
	bundles, err := service.BundleRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve bundles: %v", err)
	}
	return bundles, nil
}

func (service *BundleService) Save(bundle *model.Bundle) error {
	err := service.BundleRepo.Save(bundle)
	if err != nil {
		return err
	}
	return nil
}

func (service *BundleService) Update(bundle *model.Bundle) error {
	existingBundle, err := service.BundleRepo.Get(bundle.ID.String())
	if err != nil {
		return fmt.Errorf("failed to find bundle with ID %s: %v", bundle.ID, err)
	}

	existingBundle.ID = bundle.ID
	existingBundle.Name = bundle.Name
	existingBundle.Price = bundle.Price
	existingBundle.Status = bundle.Status
	existingBundle.Image = bundle.Image
	existingBundle.UserID = bundle.UserID
	err = service.BundleRepo.Update(&existingBundle)

	if err != nil {
		return fmt.Errorf("failed to update tour: %v", err)
	}
	return nil
}

func (service *BundleService) Delete(id string) error {
	_, err := service.BundleRepo.Get(id)
	if err != nil {
		return fmt.Errorf("failed to find bundle with ID %s: %v", id, err)
	}

	err = service.BundleRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete bundle: %v", err)
	}
	return nil
}
