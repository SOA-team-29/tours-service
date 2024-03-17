package repo

import (
	"tours/model"

	"gorm.io/gorm"
)

type TourPointRepository struct {
	DatabaseConnection *gorm.DB
}

func (tourPointRepo *TourPointRepository) CreateTourPoint(tourPoint *model.TourPoint) error {
	dbResult := tourPointRepo.DatabaseConnection.Create(tourPoint)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (tourPointRepo *TourPointRepository) GetAllPointsByTour(tourID int) ([]model.TourPoint, error) {
	var points []model.TourPoint

	if err := tourPointRepo.DatabaseConnection.Where("tour_id = ?", tourID).Find(&points).Error; err != nil {
		return nil, err
	}
	return points, nil
}
