package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"tours/service"
)

type EquipmentHandler struct {
	EquipmentService *service.EquipmentService
}

func (h *EquipmentHandler) GetAll(w http.ResponseWriter, r *http.Request) {

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

	equipments, err := h.EquipmentService.GetAll(page, pageSize)
	if err != nil {
		log.Println("Error getting equipments :", err)
		http.Error(w, "Failed to get equipments", http.StatusInternalServerError)
		return
	}

	log.Println("Equipments:", equipments)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(equipments)
}

func (h *EquipmentHandler) GetOtherEquipment(w http.ResponseWriter, r *http.Request) {
	log.Println("Stigao zahtev za ostalu opremu")
	idsStr := r.URL.Query().Get("ids")
	ids := []int{}

	if idsStr != "" {
		idStrings := strings.Split(idsStr, ",")
		for _, idStr := range idStrings {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				log.Println("Error converting ID:", err)
				http.Error(w, "Invalid ID provided", http.StatusBadRequest)
				return
			}
			ids = append(ids, id)
		}
	}
	log.Println(ids)
	equipments, err := h.EquipmentService.GetOtherEquipment(ids)
	if err != nil {
		log.Println("Error getting other equipments:", err)
		http.Error(w, "Failed to get other equipments", http.StatusInternalServerError)
		return
	}

	log.Println("Other Equipments:", equipments)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(equipments)
}

/*
func (h *EquipmentHandler) GetTouristEquipment(w http.ResponseWriter, r *http.Request) {
	log.Println("Stigao zahtev za turist opremu")
	idsStr := r.URL.Query().Get("ids")
	ids := []int{}
	log.Println(ids)
	if idsStr != "" {
		idStrings := strings.Split(idsStr, ",")
		for _, idStr := range idStrings {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				log.Println("Error converting ID:", err)
				http.Error(w, "Invalid ID provided", http.StatusBadRequest)
				return
			}
			ids = append(ids, id)
		}
	}
	log.Println(ids)
	tequipments, err := h.EquipmentService.GetTouristEquipment(ids)
	if err != nil {
		log.Println("Error getting other equipments:", err)
		http.Error(w, "Failed to get other equipments", http.StatusInternalServerError)
		return
	}

	log.Println("Tourist Equipments:", tequipments)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tequipments)
}*/

func (h *EquipmentHandler) GetTouristEquipment(w http.ResponseWriter, r *http.Request) {
	log.Println("Stigao zahtev za turist opremu")
	idsStr := r.URL.Query().Get("ids")
	log.Println(idsStr) // Ispisujemo vrednost idsStr da bismo videli šta dobijamo iz query stringa
	ids := []int{}
	log.Println(idsStr)
	if idsStr != "" {
		idStrings := strings.Split(idsStr, ",")
		for _, idStr := range idStrings {
			idStr = strings.TrimSpace(idStr) // Uklonite prazne prostore oko ID-ja
			id, err := strconv.Atoi(idStr)
			if err != nil {
				log.Println("Error converting ID:", err)
				http.Error(w, "Invalid ID provided", http.StatusBadRequest)
				return
			}
			ids = append(ids, id)
		}
	}
	log.Println(ids)
	tequipments, err := h.EquipmentService.GetTouristEquipment(ids)
	if err != nil {
		log.Println("Error getting other equipments:", err)
		http.Error(w, "Failed to get other equipments", http.StatusInternalServerError)
		return
	}

	log.Println("Tourist Equipments:", tequipments)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tequipments)
}
