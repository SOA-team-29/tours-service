package repo

import (
	"tours/model"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TouristEquipmentRepository struct {
	DatabaseConnection *gorm.DB
}

func (touristEquipmentRepo *TouristEquipmentRepository) CreateTouristEquipment(touristId int) error {
	// Kreiranje instance turističke opreme sa praznom pq.Int32Array za Equipment
	touristEquipment := model.TouristEquipment{
		TouristId: touristId,
		Equipment: pq.Int32Array{}, // Pravimo praznu pq.Int32Array
		// Postavljanje ostalih atributa prema potrebi
	}

	// Implementacija logike za čuvanje turističke opreme u bazi podataka
	if err := touristEquipmentRepo.DatabaseConnection.Create(&touristEquipment).Error; err != nil {
		return err
	}

	return nil
}

func (touristEquipmentRepo *TouristEquipmentRepository) AddToMyEquipment(touristId, equipmentId int) error {
	// Dobijanje postojeće turističke opreme na osnovu ID-ja
	touristEquipment, err := touristEquipmentRepo.GetTouristEquipment(touristId)
	if err != nil {
		return err
	}

	// Dodavanje opreme u turističku opremu
	touristEquipment.Equipment = append(touristEquipment.Equipment, int32(equipmentId))

	// Implementacija logike za ažuriranje turističke opreme u bazi podataka
	if err := touristEquipmentRepo.DatabaseConnection.Save(&touristEquipment).Error; err != nil {
		return err
	}

	return nil
}

func (touristEquipmentRepo *TouristEquipmentRepository) DeleteFromMyEquipment(touristId, equipmentId int) error {
	// Dobijanje postojeće turističke opreme na osnovu ID-ja
	touristEquipment, err := touristEquipmentRepo.GetTouristEquipment(touristId)
	if err != nil {
		return err
	}

	// Brisanje opreme iz turističke opreme
	var updatedEquipmentIDs []int32
	for _, id := range touristEquipment.Equipment {
		if int(id) != equipmentId {
			updatedEquipmentIDs = append(updatedEquipmentIDs, id)
		}
	}
	touristEquipment.Equipment = updatedEquipmentIDs

	// Implementacija logike za ažuriranje turističke opreme u bazi podataka
	if err := touristEquipmentRepo.DatabaseConnection.Save(&touristEquipment).Error; err != nil {
		return err
	}

	return nil
}

func (r *TouristEquipmentRepository) GetTouristEquipment(touristID int) (*model.TouristEquipment, error) {
	var touristEquipment model.TouristEquipment
	if err := r.DatabaseConnection.First(&touristEquipment, "tourist_id = ?", touristID).Error; err != nil {
		return nil, err
	}

	return &touristEquipment, nil
}
