package repo

import (
	"errors"
	"time"
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

	return nil
}

func (tourRepo *TourRepository) GetToursByGuideID(guideID int, page, pageSize int) ([]model.Tour, error) {
	tours := []model.Tour{}
	tourCharacteristics := []model.TourCharacteristic{}

	// Pretraga tura bez uključivanja polja tourCharacteristics
	dbResult := tourRepo.DatabaseConnection.Where("user_id = ?", guideID).Omit("tour_characteristics").Find(&tours)
	for i := range tours {
		tourRepo.DatabaseConnection.Model(&model.Tour{}).Where("id=?", tours[i].ID).Pluck("tour_characteristics", &tourCharacteristics)
		tours[i].TourCharacteristics = tourCharacteristics
	}
	if dbResult != nil {
		return tours, dbResult.Error
	}
	return tours, nil
}

func (tourRepo *TourRepository) GetTourByID(ID int) (model.Tour, error) {

	tour := model.Tour{}
	tourCharacteristics := []model.TourCharacteristic{}

	dbResult := tourRepo.DatabaseConnection.Where("id = ?", ID).Omit("tour_characteristics").First(&tour)

	tourRepo.DatabaseConnection.Model(&model.Tour{}).Where("id=?", tour.ID).Pluck("tour_characteristics", &tourCharacteristics)
	tour.TourCharacteristics = tourCharacteristics

	if dbResult != nil {
		return tour, dbResult.Error
	}
	return tour, nil
}

func (tourRepo *TourRepository) GetAllTours(page, pageSize int) ([]model.Tour, error) {

	tours := []model.Tour{}
	tourCharacteristics := []model.TourCharacteristic{}

	dbResult := tourRepo.DatabaseConnection.Omit("tour_characteristics").Find(&tours)
	for i := range tours {
		tourRepo.DatabaseConnection.Model(&model.Tour{}).Where("id=?", tours[i].ID).Pluck("tour_characteristics", &tourCharacteristics)
		tours[i].TourCharacteristics = tourCharacteristics
	}
	if dbResult != nil {
		return tours, dbResult.Error
	}
	return tours, nil
}

func (tourRepo *TourRepository) PublishTour(tourID int) error {

	tour := model.Tour{}
	tourCharacteristics := []model.TourCharacteristic{}

	tourRepo.DatabaseConnection.Omit("tour_characteristics").First(&tour, tourID)

	tourRepo.DatabaseConnection.Model(&model.Tour{}).Where("id=?", tour.ID).Pluck("tour_characteristics", &tourCharacteristics)
	tour.TourCharacteristics = tourCharacteristics
	/*
		if len(points) < 2 {
			return errors.New("tour must have at least two tour points to publish")
		}*/
	if (tourCharacteristics[len(tourCharacteristics)-1].Duration) < 0.1 {
		return errors.New("tour must have characteristic for tour to publish")
	}
	if tour.Name == "" {
		return errors.New("tour name cannot be empty")
	}
	if tour.Description == "" {
		return errors.New("tour description cannot be empty")
	}
	if len(tour.Tags) == 0 {
		return errors.New("tour must have at least one tag")
	}
	now := time.Now()
	tour.PublishedDateTime = &now

	tour.Status = model.Published

	if err := tourRepo.DatabaseConnection.Save(&tour).Error; err != nil {
		return err
	}
	return nil

}

func (tourRepo *TourRepository) ArchiveTour(tourID int) error {

	tour := model.Tour{}
	tourCharacteristics := []model.TourCharacteristic{}

	tourRepo.DatabaseConnection.Omit("tour_characteristics").First(&tour, tourID)

	tourRepo.DatabaseConnection.Model(&model.Tour{}).Where("id=?", tour.ID).Pluck("tour_characteristics", &tourCharacteristics)
	tour.TourCharacteristics = tourCharacteristics
	now := time.Now()
	tour.ArchivedDateTime = &now

	tour.Status = model.Archived

	if err := tourRepo.DatabaseConnection.Save(&tour).Error; err != nil {
		return err
	}
	return nil
}

func (tourRepo *TourRepository) SetTourCharacteristics(tourID int, characteristics []model.TourCharacteristic) error {
	var tour model.Tour
	if err := tourRepo.DatabaseConnection.First(&tour, tourID).Error; err != nil {
		return err
	}

	tour.TourCharacteristics = characteristics
	if err := tourRepo.DatabaseConnection.Save(&tour).Error; err != nil {
		return err
	}

	return nil
}
