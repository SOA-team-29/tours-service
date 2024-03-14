package model

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
