package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"tours/model"
	"tours/service"

	"github.com/gorilla/mux"
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

func (h *TourPointHandler) GetAllPointsByTour(w http.ResponseWriter, r *http.Request) {

	log.Println("Received request to get points by tour ID")

	params := mux.Vars(r)
	tourIDStr, ok := params["tourId"]
	if !ok {
		log.Println("Tour ID not provided")
		http.Error(w, "Tour ID not provided", http.StatusBadRequest)
		return
	}
	tourID, err := strconv.Atoi(tourIDStr)
	if err != nil {
		log.Println("Invalid tour ID:", err)
		http.Error(w, "Invalid tour ID", http.StatusBadRequest)
		return
	}

	points, err := h.TourPointService.GetAllPointsByTour(tourID)
	if err != nil {
		log.Println("Error getting points by tour ID:", err)
		http.Error(w, "Failed to get points by tour ID", http.StatusInternalServerError)
		return
	}
	log.Println("Points:", points)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(points)
}
