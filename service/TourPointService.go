package service

import (
	"tours/model"
	"tours/repo"
)

type TourPointService struct {
	TourPointRepo *repo.TourPointRepository
}

func (service *TourPointService) CreateTourPoint(tourPoint *model.TourPoint) error {
	err := service.TourPointRepo.CreateTourPoint(tourPoint)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourPointService) GetAllPointsByTour(tourID int) ([]model.TourPoint, error) {
	// Pozivanje odgovarajuće metode iz repozitorijuma za dobijanje tura po ID-u vodiča
	points, err := service.TourPointRepo.GetAllPointsByTour(tourID)
	if err != nil {
		return nil, err
	}
	return points, nil
}
