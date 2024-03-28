package service

import (
	"tours/model"
	"tours/repo"
)

type EquipmentService struct {
	EquipmentRepo *repo.EquipmentRepository
}

func (service *EquipmentService) GetAll(page, pageSize int) (*[]model.Equipment, error) {
	equipments, err := service.EquipmentRepo.GetAll(page, pageSize)
	if err != nil {
		return nil, err
	}
	return &equipments, nil
}

func (service *EquipmentService) GetOtherEquipment(ids []int) (*[]model.Equipment, error) {
	equipments, err := service.EquipmentRepo.GetOtherEquipment(ids)
	if err != nil {
		return nil, err
	}
	return &equipments, nil
}
func (service *EquipmentService) GetTouristEquipment(ids []int) (*[]model.Equipment, error) {
	tequipments, err := service.EquipmentRepo.GetTouristEquipment(ids)
	if err != nil {
		return nil, err
	}
	return &tequipments, nil
}
