package model

import (
	"time"

	"github.com/lib/pq"
)

type TourReview struct {
	ID             int            `json:"id"`
	Grade          float64        `json:"grade"`
	Comment        string         `json:"comment"`
	TouristID      int            `json:"touristId"`
	AttendanceDate time.Time      `json:"attendanceDate"`
	ReviewDate     time.Time      `json:"reviewDate"`
	Images         pq.StringArray `json:"images" gorm:"type:text[]"`
	TourID         int64          `json:"tourId"`
}
