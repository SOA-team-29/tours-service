package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type TransportType int

const (
	Walking TransportType = iota
	Biking
	Driving
)

type TourCharacteristic struct {
	Distance      float64       `json:"distance"`
	Duration      float64       `json:"duration"`
	TransportType TransportType `json:"transportType"`
}

func (t *TourCharacteristic) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Scan source is not []byte")
	}
	return json.Unmarshal(bytes, t)
}

// Value implements the driver.Valuer interface
func (t TourCharacteristic) Value() (driver.Value, error) {
	return json.Marshal(t)
}
