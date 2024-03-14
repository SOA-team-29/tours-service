package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TourPoint struct {
	ID          uuid.UUID `json:"id"`
	TourID      int64     `json:"tourId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	ImageURL    string    `json:"imageUrl"`
	Secret      string    `json:"secret"`
}

func (tourPoint *TourPoint) BeforeCreate(tx *gorm.DB) (err error) {
	tourPoint.ID = uuid.New()
	return nil
}
