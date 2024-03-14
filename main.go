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
	connection_url := "user=postgres password=super dbname=SOA port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(connection_url), &gorm.Config{})

	if err != nil {
		print(err)
		return nil
	}
	database.AutoMigrate(&model.Tour{}, &model.TourPoint{}, &model.TourReview{})

	return database
}

func startServer(tourHandler *handler.TourHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/tours", tourHandler.CreateTour).Methods("POST")

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

	startServer(tourHandler)
}
