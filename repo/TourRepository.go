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

func (tourRepo *TourRepository) GetTourByID(ID int) (model.Tour, error) {
	var tour model.Tour
	if err := tourRepo.DatabaseConnection.Where("id = ?", ID).First(&tour).Error; err != nil {
		return model.Tour{}, err
	}
	return tour, nil
}

func (tourRepo *TourRepository) GetAllTours(page, pageSize int) ([]model.Tour, error) {
	var tours []model.Tour

	if err := tourRepo.DatabaseConnection.Find(&tours).Error; err != nil {
		return nil, err
	}
	return tours, nil
}

func (tourRepo *TourRepository) PublishTour(tourID int) error {
	var tour model.Tour
	if err := tourRepo.DatabaseConnection.First(&tour, tourID).Error; err != nil {
		return err
	}

	tour.Status = model.Published

	if err := tourRepo.DatabaseConnection.Save(&tour).Error; err != nil {
		return err
	}

	return nil
}

func (tourRepo *TourRepository) ArchiveTour(tourID int) error {
	var tour model.Tour
	if err := tourRepo.DatabaseConnection.First(&tour, tourID).Error; err != nil {
		return err
	}

	tour.Status = model.Archived

	if err := tourRepo.DatabaseConnection.Save(&tour).Error; err != nil {
		return err
	}

	return nil
}
