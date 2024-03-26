package repo

import (
	"tours/model"

	"gorm.io/gorm"
)

type TourReviewRepository struct {
	DatabaseConnection *gorm.DB
}

func (tourReviewRepo *TourReviewRepository) CreateTourReview(tourReview *model.TourReview) error {
	dbResult := tourReviewRepo.DatabaseConnection.Create(tourReview)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (tourReviewRepo *TourReviewRepository) GetAll(page, pageSize int) ([]model.TourReview, error) {

	reviews := []model.TourReview{}

	dbResult := tourReviewRepo.DatabaseConnection.Find(&reviews)

	if dbResult != nil {
		return reviews, dbResult.Error
	}
	return reviews, nil
}
