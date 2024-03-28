package model

import "github.com/lib/pq"

type TouristEquipment struct {
	ID        int           `json:"id"`
	TouristId int           `json:"touristId"`
	Equipment pq.Int32Array ` gorm:"type:integer[]" json:"equipment"`
}
