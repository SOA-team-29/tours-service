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
