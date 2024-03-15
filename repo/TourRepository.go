package repo

import (
	"tours/model"

	"gorm.io/gorm"
)

type TourRepository struct {
	DatabaseConnection *gorm.DB
}

func (tourRepo *TourRepository) CreateTour(tour *model.Tour) error {
	dbResult := tourRepo.DatabaseConnection.Create(tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
func (tourRepo *TourRepository) GetToursByGuideID(guideID int, page, pageSize int) ([]model.Tour, error) {
	var tours []model.Tour
	offset := (page - 1) * pageSize
	if err := tourRepo.DatabaseConnection.Where("user_id = ?", guideID).Offset(offset).Limit(pageSize).Find(&tours).Error; err != nil {
		return nil, err
	}
	return tours, nil
}
