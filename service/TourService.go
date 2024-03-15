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
