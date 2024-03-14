package model

import (
	"time"

	"github.com/lib/pq"
)

type DifficultyLevel int

const (
	Easy DifficultyLevel = iota
	Moderate
	Difficult
)

type TourStatus int

const (
	Draft TourStatus = iota
	Published
	Archived
)

type Tour struct {
	ID                  int                  `json:"id"`
	Name                string               `json:"name"`
	DifficultyLevel     DifficultyLevel      `json:"difficultyLevel"`
	Description         string               `json:"description"`
	Tags                pq.StringArray       `json:"tags" gorm:"type:text[]"`
	Status              TourStatus           `json:"status"`
	Price               int                  `json:"price"`
	UserId              int                  `json:"userId"`
	PublishedDateTime   *time.Time           `json:"publishedDateTime,omitempty"`
	ArchivedDateTime    *time.Time           `json:"archivedDateTime,omitempty"`
	TourPoints          []TourPoint          `json:"tourPoints"`
	TourCharacteristics []TourCharacteristic `gorm:"type:TourCharacteristic[]"`
	TourReviews         []TourReview         `json:"tourReviews"`
}
