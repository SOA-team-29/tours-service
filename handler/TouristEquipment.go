package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"tours/service"

	"github.com/gorilla/mux"
)

type TouristEquipmentHandler struct {
	TouristEquipmentService *service.TouristEquipmentService
}

func (h *TouristEquipmentHandler) CreateTouristEquipment(w http.ResponseWriter, r *http.Request) {
	log.Println("Stigao zahtev za kreiranje turističke opreme")

	// Uzimanje id iz URL putanje
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Error converting ID:", err)
		http.Error(w, "Invalid ID provided", http.StatusBadRequest)
		return
	}

	// Pozivanje odgovarajuće metode iz servisa
	result := h.TouristEquipmentService.CreateTouristEquipment(id)

	// Vraćanje odgovora
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *TouristEquipmentHandler) AddToMyEquipment(w http.ResponseWriter, r *http.Request) {
	log.Println("Stigao zahtev za dodavanje u moju opremu")

	// Uzimanje touristId i equipmentId iz URL putanje
	vars := mux.Vars(r)
	touristIdStr := vars["touristId"]
	touristId, err := strconv.Atoi(touristIdStr)
	if err != nil {
		log.Println("Error converting tourist ID:", err)
		http.Error(w, "Invalid tourist ID provided", http.StatusBadRequest)
		return
	}
	equipmentIdStr := vars["equipmentId"]
	equipmentId, err := strconv.Atoi(equipmentIdStr)
	if err != nil {
		log.Println("Error converting equipment ID:", err)
		http.Error(w, "Invalid equipment ID provided", http.StatusBadRequest)
		return
	}

	// Pozivanje odgovarajuće metode iz servisa
	odg := h.TouristEquipmentService.AddToMyEquipment(touristId, equipmentId)

	// Vraćanje odgovora
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(odg)
}

func (h *TouristEquipmentHandler) DeleteFromMyEquipment(w http.ResponseWriter, r *http.Request) {
	log.Println("Stigao zahtev za brisanje iz moje opreme")

	// Uzimanje touristId i equipmentId iz URL putanje
	vars := mux.Vars(r)
	touristIdStr := vars["touristId"]
	touristId, err := strconv.Atoi(touristIdStr)
	if err != nil {
		log.Println("Error converting tourist ID:", err)
		http.Error(w, "Invalid tourist ID provided", http.StatusBadRequest)
		return
	}
	equipmentIdStr := vars["equipmentId"]
	equipmentId, err := strconv.Atoi(equipmentIdStr)
	if err != nil {
		log.Println("Error converting equipment ID:", err)
		http.Error(w, "Invalid equipment ID provided", http.StatusBadRequest)
		return
	}

	odg := h.TouristEquipmentService.DeleteFromMyEquipment(touristId, equipmentId)

	// Vraćanje odgovora
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(odg)
}

func (h *TouristEquipmentHandler) GetTouristEquipment(w http.ResponseWriter, r *http.Request) {
	log.Println("USAO")
	vars := mux.Vars(r)
	touristID, err := strconv.Atoi(vars["touristId"])
	if err != nil {
		http.Error(w, "Invalid tourist ID provided", http.StatusBadRequest)
		return
	}

	touristEquipment, err := h.TouristEquipmentService.GetTouristEquipment(touristID)
	log.Println(touristEquipment)
	if err != nil {
		log.Println("Error getting tourist equipment:", err)
		http.Error(w, "Failed to get tourist equipment", http.StatusInternalServerError)
		return
	}

	// Pretvaranje turističke opreme u JSON format i slanje kao odgovor
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(touristEquipment)
}
