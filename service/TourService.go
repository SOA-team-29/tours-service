package service

import (
	"tours/model"
	"tours/repo"
)

type TourService struct {
	TourRepo *repo.TourRepository
}

func (service *TourService) CreateTour(tour *model.Tour) error {
	err := service.TourRepo.CreateTour(tour)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourService) GetToursByGuideID(guideID int, page, pageSize int) ([]model.Tour, error) {
	// Pozivanje odgovarajuće metode iz repozitorijuma za dobijanje tura po ID-u vodiča
	tours, err := service.TourRepo.GetToursByGuideID(guideID, page, pageSize)
	if err != nil {
		return nil, err
	}
	return tours, nil
}

func (service *TourService) GetTourByID(ID int) (model.Tour, error) {
	tour, err := service.TourRepo.GetTourByID(ID)
	if err != nil {
		return model.Tour{}, err
	}
	return tour, nil
}

func (service *TourService) GetAllTours(page, pageSize int) ([]model.Tour, error) {
	tours, err := service.TourRepo.GetAllTours(page, pageSize)
	if err != nil {
		return nil, err
	}
	return tours, nil
}

func (service *TourService) PublishTour(tourID int) error {

	err := service.TourRepo.PublishTour(tourID)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourService) ArchiveTour(tourID int) error {
	err := service.TourRepo.ArchiveTour(tourID)
	if err != nil {
		return err
	}
	return nil
}
