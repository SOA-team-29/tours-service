package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"tours/model"
	"tours/service"
)

type TourReviewHandler struct {
	TourReviewService *service.TourReviewService
}

func (tourReviewHandler *TourReviewHandler) CreateTourReview(writer http.ResponseWriter, req *http.Request) {
	var tourReview model.TourReview
	err := json.NewDecoder(req.Body).Decode(&tourReview)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = tourReviewHandler.TourReviewService.CreateTourReview(&tourReview)
	if err != nil {
		println("Error while creating a new tourReview")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (h *TourReviewHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	pageSizeStr := r.URL.Query().Get("pageSize")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}

	reviews, err := h.TourReviewService.GetAll(page, pageSize)
	if err != nil {
		log.Println("Error getting reviews :", err)
		http.Error(w, "Failed to get reviews", http.StatusInternalServerError)
		return
	}

	log.Println("Reviews:", reviews)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}
