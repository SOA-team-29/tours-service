package model

type TourPoint struct {
	ID          int     `json:"id"`
	TourID      int64   `json:"tourId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	ImageURL    string  `json:"imageUrl"`
	Secret      string  `json:"secret"`
}
