package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"tours/model"
	"tours/service"
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
