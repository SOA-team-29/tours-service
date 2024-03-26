package service

import (
	"tours/model"
	"tours/repo"
)

type TourReviewService struct {
	TourReviewRepo *repo.TourReviewRepository
}

func (service *TourReviewService) CreateTourReview(tourReview *model.TourReview) error {
	err := service.TourReviewRepo.CreateTourReview(tourReview)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourReviewService) GetAll(page, pageSize int) (*[]model.TourReview, error) {
	reviews, err := service.TourReviewRepo.GetAll(page, pageSize)
	if err != nil {
		return nil, err
	}
	return &reviews, nil
}
