package main

import (
	"log"
	"net/http"
	"tours/handler"
	"tours/model"
	"tours/repo"
	"tours/service"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connectionURL := "user=postgres password=super dbname=SOA host=database port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(connectionURL), &gorm.Config{})

	if err != nil {
		print(err)
		return nil
	}
	database.AutoMigrate(&model.Tour{}, &model.TourPoint{}, &model.TourReview{}, &model.Equipment{}, &model.TouristEquipment{})

	return database
}

func startServer(tourHandler *handler.TourHandler, tourPointHandler *handler.TourPointHandler, tourReviewHandler *handler.TourReviewHandler,
	equipmentHandler *handler.EquipmentHandler, touristEquipmentHandler *handler.TouristEquipmentHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/tours", tourHandler.CreateTour).Methods("POST")
	router.HandleFunc("/tours/see/all", tourHandler.GetAllTours).Methods("GET")
	router.HandleFunc("/tours/{id}", tourHandler.GetTourByID).Methods("GET")
	router.HandleFunc("/toursByGuideId/{userId}", tourHandler.GetToursByGuideID).Methods("GET")
	router.HandleFunc("/tours/publish/{tourId}", tourHandler.PublishTour).Methods("PUT")
	router.HandleFunc("/tours/archive/{id}", tourHandler.ArchiveTour).Methods("PUT")
	router.HandleFunc("/tours/characteristics/{tourId}", tourHandler.SetTourCharacteristic).Methods("PUT")
	router.HandleFunc("/tourPoint", tourPointHandler.CreateTourPoint).Methods("POST")
	router.HandleFunc("/tourPoint/allPointsInTour/{tourId}", tourPointHandler.GetAllPointsByTour).Methods("GET")
	router.HandleFunc("/tourReviews/create", tourReviewHandler.CreateTourReview).Methods("POST")
	router.HandleFunc("/tourReviews/see", tourReviewHandler.GetAll).Methods("GET")
	router.HandleFunc("/equipment/getOtherEquipment", equipmentHandler.GetOtherEquipment).Methods("GET")
	router.HandleFunc("/equipment/getAll", equipmentHandler.GetAll).Methods("GET")
	router.HandleFunc("/equipment/tourist/getEquipment/", equipmentHandler.GetTouristEquipment).Methods("GET")

	router.HandleFunc("/getTouristEquipment/{touristId}", touristEquipmentHandler.GetTouristEquipment).Methods("GET")
	router.HandleFunc("/touristEquipment/createTouristEquipment/{id}", touristEquipmentHandler.CreateTouristEquipment).Methods("POST")
	router.HandleFunc("/touristEquipment/addToMyEquipment/{touristId}/{equipmentId}", touristEquipmentHandler.AddToMyEquipment).Methods("PUT")
	router.HandleFunc("/touristEquipment/deleteFromMyEquipment/{touristId}/{equipmentId}", touristEquipmentHandler.DeleteFromMyEquipment).Methods("PUT")

	println("Server starting")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	tourRepo := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{TourRepo: tourRepo}
	tourHandler := &handler.TourHandler{TourService: tourService}

	tourPointRepo := &repo.TourPointRepository{DatabaseConnection: database}
	tourPointService := &service.TourPointService{TourPointRepo: tourPointRepo}
	tourPointHandler := &handler.TourPointHandler{TourPointService: tourPointService}

	tourReviewRepo := &repo.TourReviewRepository{DatabaseConnection: database}
	tourReviewService := &service.TourReviewService{TourReviewRepo: tourReviewRepo}
	tourReviewHandler := &handler.TourReviewHandler{TourReviewService: tourReviewService}

	equipmentRepo := &repo.EquipmentRepository{DatabaseConnection: database}
	equipmentService := &service.EquipmentService{EquipmentRepo: equipmentRepo}
	equipmentHandler := &handler.EquipmentHandler{EquipmentService: equipmentService}

	touristEquipmentRepo := &repo.TouristEquipmentRepository{DatabaseConnection: database}
	touristEquipmentService := &service.TouristEquipmentService{TouristEquipmentRepo: touristEquipmentRepo}
	touristEquipmentHandler := &handler.TouristEquipmentHandler{TouristEquipmentService: touristEquipmentService}

	startServer(tourHandler, tourPointHandler, tourReviewHandler, equipmentHandler, touristEquipmentHandler)
}
