package service

import (
	"log"
	"tours/model"
	"tours/repo"
)

type TouristEquipmentService struct {
	TouristEquipmentRepo *repo.TouristEquipmentRepository
}

func (touristequipmentService *TouristEquipmentService) CreateTouristEquipment(id int) error {
	log.Println("Kreiranje turističke opreme")

	// Pozivanje odgovarajuće metode iz repozitorijuma
	if err := touristequipmentService.TouristEquipmentRepo.CreateTouristEquipment(id); err != nil {
		return err
	}

	return nil
}

func (touristequipmentService *TouristEquipmentService) AddToMyEquipment(touristId, equipmentId int) *model.TouristEquipment {
	log.Println("Dodavanje u moju opremu")

	// Pozivanje metode za dodavanje opreme u turističku opremu iz repozitorijuma
	if err := touristequipmentService.TouristEquipmentRepo.AddToMyEquipment(touristId, equipmentId); err != nil {
		return nil
	}

	// Dobijanje ažurirane turističke opreme
	updatedTouristEquipment, err := touristequipmentService.TouristEquipmentRepo.GetTouristEquipment(touristId)
	if err != nil {
		return nil
	}

	return updatedTouristEquipment
}

func (touristequipmentService *TouristEquipmentService) DeleteFromMyEquipment(touristId, equipmentId int) *model.TouristEquipment {
	log.Println("Brisanje iz moje opreme")

	// Pozivanje metode za dodavanje opreme u turističku opremu iz repozitorijuma
	if err := touristequipmentService.TouristEquipmentRepo.DeleteFromMyEquipment(touristId, equipmentId); err != nil {
		return nil
	}

	// Dobijanje ažurirane turističke opreme
	updatedTouristEquipment, err := touristequipmentService.TouristEquipmentRepo.GetTouristEquipment(touristId)
	if err != nil {
		return nil
	}

	return updatedTouristEquipment
}

func (s *TouristEquipmentService) GetTouristEquipment(touristID int) (*model.TouristEquipment, error) {
	touristEquipment, err := s.TouristEquipmentRepo.GetTouristEquipment(touristID)
	if err != nil {
		return nil, err
	}
	return touristEquipment, nil
}
