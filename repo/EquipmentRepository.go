package repo

import (
	"log"
	"tours/model"

	"gorm.io/gorm"
)

type EquipmentRepository struct {
	DatabaseConnection *gorm.DB
}

func (equipmentRepo *EquipmentRepository) GetAll(page, pageSize int) ([]model.Equipment, error) {

	equipments := []model.Equipment{}

	dbResult := equipmentRepo.DatabaseConnection.Find(&equipments)

	if dbResult != nil {
		return equipments, dbResult.Error
	}
	return equipments, nil
}
func (equipmentRepo *EquipmentRepository) GetOtherEquipment(ids []int) ([]model.Equipment, error) {

	allEquipments := []model.Equipment{}
	if err := equipmentRepo.DatabaseConnection.Find(&allEquipments).Error; err != nil {
		return nil, err
	}

	// Kopiranje svih oprema u rezultujuću listu
	result := make([]model.Equipment, len(allEquipments))
	copy(result, allEquipments)

	// Uklanjanje opreme sa zadatim ID-jevima
	for _, id := range ids {
		for i, equipment := range result {
			if equipment.ID == id {
				// Uklanjanje opreme sa zadatim ID-jevima
				result = append(result[:i], result[i+1:]...)
				break
			}
		}
	}

	return result, nil
}
func (equipmentRepo *EquipmentRepository) GetTouristEquipment(ids []int) ([]model.Equipment, error) {
	result := []model.Equipment{}
	log.Println(ids)
	for _, id := range ids {
		var equipment model.Equipment
		if err := equipmentRepo.DatabaseConnection.Where("id = ?", id).First(&equipment).Error; err != nil {
			return nil, err
		}
		result = append(result, equipment)
	}

	return result, nil
}
