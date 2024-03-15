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
