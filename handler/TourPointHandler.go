package handler

import (
	"encoding/json"
	"net/http"
	"tours/model"
	"tours/service"
)

type TourPointHandler struct {
	TourPointService *service.TourPointService
}

func (tourPointHandler *TourPointHandler) CreateTourPoint(writer http.ResponseWriter, req *http.Request) {
	var tourPoint model.TourPoint
	err := json.NewDecoder(req.Body).Decode(&tourPoint)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = tourPointHandler.TourPointService.CreateTourPoint(&tourPoint)
	if err != nil {
		println("Error while creating a new tourPoint")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
