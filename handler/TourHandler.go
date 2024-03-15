package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"tours/model"
	"tours/service"

	"github.com/gorilla/mux"
)

type TourHandler struct {
	TourService *service.TourService
}

func (h *TourHandler) CreateTour(w http.ResponseWriter, r *http.Request) {
	// Ispisivanje podataka o zahtevu koji dolazi
	log.Println("Received request to create tour")

	// Čitanje JSON podataka iz tela zahteva
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Ispisivanje JSON podataka pre dekodiranja
	log.Println("Received JSON data:", string(body))

	// Modifikacija polja difficultyLevel i status pre dekodiranja
	modifiedBody := modifyJSON(body)

	// Dekodiranje JSON podataka iz tela zahteva u tour objekat
	var tour model.Tour
	decoder := json.NewDecoder(bytes.NewReader(modifiedBody))

	err = decoder.Decode(&tour)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Failed to decode JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Prosleđivanje tour objekta servisu za kreiranje ture
	err = h.TourService.CreateTour(&tour)
	if err != nil {
		log.Println("Error while creating a new tour:", err)
		http.Error(w, "Failed to create a new tour", http.StatusInternalServerError)
		return
	}

	// Slanje odgovora da je tura uspešno kreirana
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

// Funkcija za modifikaciju JSON podataka
func modifyJSON(data []byte) []byte {
	var modifiedData map[string]interface{}
	if err := json.Unmarshal(data, &modifiedData); err != nil {
		log.Println("Error decoding JSON:", err)
		return data
	}

	// Konverzija difficultyLevel iz stringa u broj
	if difficulty, ok := modifiedData["difficultyLevel"].(string); ok {
		modifiedData["difficultyLevel"] = convertDifficultyToNumber(difficulty)
	}

	// Konverzija statusa iz stringa u broj
	if status, ok := modifiedData["status"].(string); ok {
		modifiedData["status"] = convertStatusToNumber(status)
	}

	// Konverzija nazad u JSON
	modifiedBody, err := json.Marshal(modifiedData)
	if err != nil {
		log.Println("Error encoding modified JSON:", err)
		return data
	}

	return modifiedBody
}

// Funkcija za konverziju difficultyLevel u broj
func convertDifficultyToNumber(difficulty string) int {
	switch difficulty {
	case "Easy":
		return 0
	case "Moderate":
		return 1
	case "Difficult":
		return 2
	default:
		return -1
	}
}

// Funkcija za konverziju statusa u broj
func convertStatusToNumber(status string) int {
	switch status {
	case "Draft":
		return 0
	case "Published":
		return 1
	case "Archived":
		return 2
	default:
		return -1
	}
}

func (h *TourHandler) GetToursByGuideID(w http.ResponseWriter, r *http.Request) {
	// Ispisivanje podataka o zahtevu koji dolazi
	log.Println("Received request to get tours by guide ID")

	// Čitanje ID vodiča iz URL parametra
	params := mux.Vars(r)
	guideIDStr, ok := params["userId"]
	if !ok {
		log.Println("Guide ID not provided")
		http.Error(w, "Guide ID not provided", http.StatusBadRequest)
		return
	}
	guideID, err := strconv.Atoi(guideIDStr)
	if err != nil {
		log.Println("Invalid guide ID:", err)
		http.Error(w, "Invalid guide ID", http.StatusBadRequest)
		return
	}

	// Čitanje stranice i veličine stranice iz URL upita (query parameters)
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1 // Ako nije navedena stranica, podrazumevano je prva stranica
	}

	pageSizeStr := r.URL.Query().Get("pageSize")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10 // Ako nije navedena veličina stranice, podrazumevano je 10
	}

	// Pozivanje odgovarajuće funkcije u servisu za dobijanje tura po ID-u vodiča
	tours, err := h.TourService.GetToursByGuideID(guideID, page, pageSize)
	if err != nil {
		log.Println("Error getting tours by guide ID:", err)
		http.Error(w, "Failed to get tours by guide ID", http.StatusInternalServerError)
		return
	}

	//Moja provera da li je nasao dobro iz baze
	log.Println("Tours:", tours)

	modifiedTours := make([]map[string]interface{}, len(tours))
	for i, tour := range tours {
		modifiedTour := map[string]interface{}{
			"id":                  tour.ID,
			"name":                tour.Name,
			"difficultyLevel":     convertDifficultyToString(int(tour.DifficultyLevel)),
			"description":         tour.Description,
			"tags":                tour.Tags,
			"status":              convertStatusToString(int(tour.Status)),
			"price":               tour.Price,
			"userId":              tour.UserId,
			"publishedDateTime":   tour.PublishedDateTime,
			"archivedDateTime":    tour.ArchivedDateTime,
			"tourPoints":          tour.TourPoints,
			"tourCharacteristics": tour.TourCharacteristics,
			"tourReviews":         tour.TourReviews,
		}
		modifiedTours[i] = modifiedTour
	}

	// Slanje odgovora sa tura podacima kao JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(modifiedTours)
}

// Funkcija za konverziju difficultyLevel u string
func convertDifficultyToString(difficulty int) string {
	switch difficulty {
	case 0:
		return "Easy"
	case 1:
		return "Moderate"
	case 2:
		return "Difficult"
	default:
		return ""
	}
}

// Funkcija za konverziju statusa u string
func convertStatusToString(status int) string {
	switch status {
	case 0:
		return "Draft"
	case 1:
		return "Published"
	case 2:
		return "Archived"
	default:
		return ""
	}
}
